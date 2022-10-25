package admin

import (
	"context"

	orgv1 "github.com/confluentinc/cc-structs/kafka/org/v1"
	"github.com/spf13/cobra"

	"github.com/confluentinc/cli/internal/pkg/utils"
)

func (c *command) newAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add <code>",
		Short: "Add a new promo code.",
		Args:  cobra.ExactArgs(1),
		RunE:  c.add,
	}
}

func (c *command) add(cmd *cobra.Command, args []string) error {
	org := &orgv1.Organization{Id: c.Context.GetOrganization().GetId()}

	if _, err := c.Client.Billing.ClaimPromoCode(context.Background(), org, args[0]); err != nil {
		return err
	}

	utils.Println(cmd, "Your promo code was successfully added.")
	return nil
}
