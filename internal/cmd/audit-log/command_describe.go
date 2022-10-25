package auditlog

import (
	"context"

	"github.com/spf13/cobra"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	v1 "github.com/confluentinc/cli/internal/pkg/config/v1"
	"github.com/confluentinc/cli/internal/pkg/errors"
	"github.com/confluentinc/cli/internal/pkg/output"
)

var (
	listFields    = []string{"ClusterId", "EnvironmentId", "ServiceAccountId", "TopicName"}
	humanLabelMap = map[string]string{
		"ClusterId":        "Cluster",
		"EnvironmentId":    "Environment",
		"ServiceAccountId": "Service Account",
		"TopicName":        "Topic Name",
	}
	structuredLabelMap = map[string]string{
		"ClusterId":        "cluster_id",
		"EnvironmentId":    "environment_id",
		"ServiceAccountId": "service_account_id",
		"TopicName":        "topic_name",
	}
)

type describeCmd struct {
	*pcmd.AuthenticatedCLICommand
}

type auditLogStruct struct {
	ClusterId        string
	EnvironmentId    string
	ServiceAccountId string
	TopicName        string
}

func newDescribeCommand(prerunner pcmd.PreRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "describe",
		Short:       "Describe the audit log configuration for this organization.",
		Args:        cobra.NoArgs,
		Annotations: map[string]string{pcmd.RunRequirement: pcmd.RequireCloudLogin},
	}

	c := &describeCmd{pcmd.NewAuthenticatedCLICommand(cmd, prerunner)}
	c.RunE = c.describe

	pcmd.AddOutputFlag(c.Command)

	return c.Command
}

func (c describeCmd) describe(cmd *cobra.Command, _ []string) error {
	if v1.GetAuditLog(c.Context.Context) == nil {
		return errors.New(errors.AuditLogsNotEnabledErrorMsg)
	}

	auditLog := c.Context.GetOrganization().GetAuditLog()

	serviceAccount, err := c.Client.User.GetServiceAccount(context.Background(), auditLog.GetServiceAccountId())
	if err != nil {
		return err
	}

	return output.DescribeObject(cmd, &auditLogStruct{
		ClusterId:        auditLog.GetClusterId(),
		EnvironmentId:    auditLog.GetAccountId(),
		ServiceAccountId: serviceAccount.GetResourceId(),
		TopicName:        auditLog.GetTopicName(),
	}, listFields, humanLabelMap, structuredLabelMap)
}
