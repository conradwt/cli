package byok

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/spf13/cobra"

	byokv1 "github.com/confluentinc/ccloud-sdk-go-v2/byok/v1"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/errors"
	"github.com/confluentinc/cli/internal/pkg/examples"
)

var encryptionKeyPolicyAws = template.Must(template.New("encryptionKeyPolicyAws").Parse(`{
	"Sid" : "Allow Confluent accounts to use the key",
	"Effect" : "Allow",
	"Principal" : {
		"AWS" : [{{range $i, $e := .}}{{if $i}},{{end}}
			"{{$e}}"{{end}}
		]
	},
	"Action" : [ "kms:Encrypt", "kms:Decrypt", "kms:ReEncrypt*", "kms:GenerateDataKey*", "kms:DescribeKey" ],
	"Resource" : "*"
}, {
	"Sid" : "Allow Confluent accounts to attach persistent resources",
	"Effect" : "Allow",
	"Principal" : {
		"AWS" : [{{range $i, $e := .}}{{if $i}},{{end}}
			"{{$e}}"{{end}}
		]
	},
	"Action" : [ "kms:CreateGrant", "kms:ListGrants", "kms:RevokeGrant" ],
	"Resource" : "*"
}`))

const (
	keyVaultCryptoServiceEncryptionUser = "e147488a-f6f5-4113-8e2d-b22465e65bf6"
	keyVaultReader                      = "21090545-7ca7-4776-b22c-e363652d74d2"
)

func (c *command) newCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <key>",
		Short: "Register a self-managed encryption key.",
		Long:  "Bring your own key to Confluent Cloud for data at rest encryption (AWS and Azure only).",
		Args:  cobra.ExactArgs(1),
		RunE:  c.create,
		Example: examples.BuildExampleString(
			examples.Example{
				Text: "Register a new self-managed encryption key for AWS:",
				Code: `confluent byok create "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`,
			},
			examples.Example{
				Text: "Register a new self-managed encryption key for Azure:",
				Code: `confluent byok create "https://vault-name.vault.azure.net/keys/key-name" --tenant "00000000-0000-0000-0000-000000000000" --key-vault "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup-name/providers/Microsoft.KeyVault/vaults/vault-name"`,
			},
		),
	}

	cmd.Flags().String("key-vault", "", "The ID of the Azure Key Vault where the key is stored.")
	cmd.Flags().String("tenant", "", "The ID of the Azure Active Directory tenant that the key vault belongs to.")
	pcmd.AddOutputFlag(cmd)

	cmd.MarkFlagsRequiredTogether("key-vault", "tenant")

	return cmd
}

func (c *command) createAwsKeyRequest(keyArn string) *byokv1.ByokV1Key {
	return &byokv1.ByokV1Key{
		Key: &byokv1.ByokV1KeyKeyOneOf{
			ByokV1AwsKey: &byokv1.ByokV1AwsKey{
				KeyArn: keyArn,
				Kind:   "AwsKey",
			},
		},
	}
}

func (c *command) createAzureKeyRequest(cmd *cobra.Command, keyString string) (*byokv1.ByokV1Key, error) {
	keyVault, err := cmd.Flags().GetString("key-vault")
	if err != nil {
		return nil, err
	}
	tenant, err := cmd.Flags().GetString("tenant")
	if err != nil {
		return nil, err
	}

	keyReq := byokv1.ByokV1Key{
		Key: &byokv1.ByokV1KeyKeyOneOf{
			ByokV1AzureKey: &byokv1.ByokV1AzureKey{
				KeyId:      keyString,
				KeyVaultId: keyVault,
				TenantId:   tenant,
				Kind:       "AzureKey",
			},
		},
	}

	return &keyReq, nil
}

func (c *command) create(cmd *cobra.Command, args []string) error {
	keyString := args[0]
	var err error
	var keyReq *byokv1.ByokV1Key

	if cmd.Flags().Changed("key-vault") && cmd.Flags().Changed("tenant") {
		keyString = removeKeyVersionFromAzureKeyId(keyString)

		keyReq, err = c.createAzureKeyRequest(cmd, keyString)
		if err != nil {
			return err
		}
	} else if isAWSKey(keyString) {
		keyReq = c.createAwsKeyRequest(keyString)
	} else {
		return errors.New(fmt.Sprintf("invalid key format: %s", keyString))
	}

	key, httpResp, err := c.V2Client.CreateByokKey(*keyReq)
	if err != nil {
		return errors.CatchCCloudV2Error(err, httpResp)
	}

	return c.outputByokKeyDescription(cmd, &key)
}

func isAWSKey(key string) bool {
	keyArn, err := arn.Parse(key)
	if err != nil {
		return false
	}

	return keyArn.Service == "kms" && strings.HasPrefix(keyArn.Resource, "key/")
}

func getPolicyCommand(key *byokv1.ByokV1Key) (string, error) {
	switch {
	case key.Key.ByokV1AwsKey != nil:
		return renderAWSEncryptionPolicy(key.Key.ByokV1AwsKey.GetRoles())
	case key.Key.ByokV1AzureKey != nil:
		return renderAzureEncryptionPolicy(key)
	default:
		return "", nil
	}
}

func renderAWSEncryptionPolicy(roles []string) (string, error) {
	buf := new(bytes.Buffer)
	if err := encryptionKeyPolicyAws.Execute(buf, roles); err != nil {
		return "", errors.New(errors.FailedToRenderKeyPolicyErrorMsg)
	}
	return buf.String(), nil
}

func renderAzureEncryptionPolicy(key *byokv1.ByokV1Key) (string, error) {
	objectId := fmt.Sprintf(`$(az ad sp show --id "%s" --query id --out tsv 2>/dev/null || az ad sp create --id "%s" --query id --out tsv)`, key.Key.ByokV1AzureKey.GetApplicationId(), key.Key.ByokV1AzureKey.GetApplicationId())

	regex := regexp.MustCompile(`^https://([^/.]+).vault.azure.net`)
	matches := regex.FindStringSubmatch(key.Key.ByokV1AzureKey.KeyId)
	if matches == nil {
		return "", errors.New(errors.FailedToRenderKeyPolicyErrorMsg)
	}

	vaultName := matches[1]

	az := []string{
		"az role assignment create \\",
		fmt.Sprintf("    --role \"%s\" \\", keyVaultCryptoServiceEncryptionUser),
		fmt.Sprintf("    --scope \"$(az keyvault show --name \"%s\" --query id --output tsv)\" \\", vaultName),
		fmt.Sprintf("    --assignee-object-id \"%s\" \\", objectId),
		"    --assignee-principal-type ServicePrincipal && \\",
		"az role assignment create \\",
		fmt.Sprintf("    --role \"%s\" \\", keyVaultReader),
		fmt.Sprintf("    --scope \"$(az keyvault show --name \"%s\" --query id --output tsv)\" \\", vaultName),
		fmt.Sprintf("    --assignee-object-id \"%s\" \\", objectId),
		"    --assignee-principal-type ServicePrincipal",
	}

	return strings.Join(az, "\n"), nil
}

func getPostCreateStepInstruction(key *byokv1.ByokV1Key) string {
	switch {
	case key.Key.ByokV1AwsKey != nil:
		return errors.CopyByokAwsPermissionsHeaderMsg
	case key.Key.ByokV1AzureKey != nil:
		return errors.RunByokAzurePermissionsHeaderMsg
	default:
		return ""
	}
}

// Best effort to remove the key version from the Azure Key ID if it is present
// For any errors, return the original key ID as is
// All further validation of the key ID is done by the BYOK API
func removeKeyVersionFromAzureKeyId(keyId string) string {
	path, err := url.Parse(keyId)
	if err != nil || len(strings.Split(path.Path, "/")) != 4 {
		return keyId
	}

	pathSegments := strings.Split(path.Path, "/")
	return keyId[:len(keyId)-len(pathSegments[3])-1]
}
