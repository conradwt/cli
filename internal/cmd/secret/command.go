package secret

import (
	"github.com/spf13/cobra"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	"github.com/confluentinc/cli/internal/pkg/secret"
)

type command struct {
	*pcmd.AuthenticatedStateFlagCommand
	flagResolver pcmd.FlagResolver
	plugin       secret.PasswordProtection
}

func New(prerunner pcmd.PreRunner, flagResolver pcmd.FlagResolver, plugin secret.PasswordProtection) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "secret",
		Short:       "Manage secrets for Confluent Platform.",
		Annotations: map[string]string{pcmd.RunRequirement: pcmd.RequireOnPremLogin},
	}

	c := &command{
		AuthenticatedStateFlagCommand: pcmd.NewAuthenticatedWithMDSStateFlagCommand(cmd, prerunner),
		flagResolver:                  flagResolver,
		plugin:                        plugin,
	}

	c.AddCommand(c.newMasterKeyCommand())
	c.AddCommand(c.newFileCommand())

	return c.Command
}
