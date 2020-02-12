package connector

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	connectv1 "github.com/confluentinc/ccloudapis/connect/v1"
	"github.com/confluentinc/go-printer"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
	v2 "github.com/confluentinc/cli/internal/pkg/config/v2"
	"github.com/confluentinc/cli/internal/pkg/errors"
	"github.com/confluentinc/cli/internal/pkg/output"
)

type command struct {
	*pcmd.AuthenticatedCLICommand
}

type describeDisplay struct {
	Name   string
	ID     string
	Status string
	Type   string
}

var (
	describeRenames      = map[string]string{}
	listFields           = []string{"ID", "Name", "Status", "Type"}
	listStructuredLabels = []string{"id", "name", "status", "type"}
)

// New returns the default command object for interacting with Connect.
func New(prerunner pcmd.PreRunner, config *v2.Config) *cobra.Command {
	cmd := &command{
		AuthenticatedCLICommand: pcmd.NewAuthenticatedCLICommand(
			&cobra.Command{
				Use:   "connector",
				Short: "Manage Kafka Connect.",
			}, config, prerunner),
	}
	cmd.init()
	return cmd.Command
}

func (c *command) init() {
	cmd := &cobra.Command{
		Use:   "describe <id>",
		Short: "Describe a connector.",
		Example: FormatDescription(`
Describe connector and task level details of a connector in the current or specified Kafka cluster context.

::

        {{.CLIName}} connector describe <id>
        {{.CLIName}} connector describe <id> --cluster <cluster-id>		`, c.Config.CLIName),
		RunE: c.describe,
		Args: cobra.ExactArgs(1),
	}
	cmd.Flags().String("cluster", "", "Kafka cluster ID.")
	cmd.Flags().SortFlags = false
	c.AddCommand(cmd)

	cmd = &cobra.Command{
		Use:   "list",
		Short: "List connectors.",
		Example: FormatDescription(`
List connectors in the current or specified Kafka cluster context.

::

        {{.CLIName}} connector list
        {{.CLIName}} connector list --cluster <cluster-id>		`, c.Config.CLIName),
		RunE: c.list,
		Args: cobra.NoArgs,
	}
	cmd.Flags().String("cluster", "", "Kafka cluster ID.")
	cmd.Flags().StringP(output.FlagName, output.ShortHandFlag, output.DefaultValue, output.Usage)
	cmd.Flags().SortFlags = false
	c.AddCommand(cmd)

	cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a connector.",
		Example: FormatDescription(`
Create connector in the current or specified Kafka cluster context.

::

        {{.CLIName}} connector create --config <file>
        {{.CLIName}} connector create --cluster <cluster-id> --config <file>		`, c.Config.CLIName),
		RunE: c.create,
		Args: cobra.NoArgs,
	}
	cmd.Flags().String("config", "", "JSON connector config file.")
	cmd.Flags().String("cluster", "", "Kafka cluster ID.")
	panicOnError(cmd.MarkFlagRequired("config"))
	cmd.Flags().SortFlags = false
	c.AddCommand(cmd)

	cmd = &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a connector.",
		Example: FormatDescription(`
Delete connector in the current or specified Kafka cluster context.

::

        {{.CLIName}} connector delete <id>
        {{.CLIName}} connector delete <id> --cluster <cluster-id>	`, c.Config.CLIName),
		RunE: c.delete,
		Args: cobra.ExactArgs(1),
	}
	cmd.Flags().String("cluster", "", "Kafka cluster ID.")
	cmd.Flags().SortFlags = false
	c.AddCommand(cmd)

	cmd = &cobra.Command{
		Use:   "update <id>",
		Short: "Update connector configuration.",
		RunE:  c.update,
		Args:  cobra.ExactArgs(1),
	}
	cmd.Flags().String("config", "", "JSON connector config file.")
	cmd.Flags().String("cluster", "", "Kafka cluster ID.")
	panicOnError(cmd.MarkFlagRequired("config"))
	cmd.Flags().SortFlags = false
	c.AddCommand(cmd)

	cmd = &cobra.Command{
		Use:   "pause <id>",
		Short: "Pause a connector.",
		Example: FormatDescription(`
Pause connector in the current or specified Kafka cluster context.

::

        {{.CLIName}} connector pause <connector-id>
        {{.CLIName}} connector pause <connector-id> --cluster <cluster-id>	`, c.Config.CLIName),
		RunE: c.pause,
		Args: cobra.ExactArgs(1),
	}
	cmd.Flags().String("cluster", "", "Kafka cluster ID.")
	cmd.Flags().SortFlags = false
	c.AddCommand(cmd)

	cmd = &cobra.Command{
		Use:   "resume <id>",
		Short: "Resume a connector.",
		Example: FormatDescription(`
Resume connector in the current or specified Kafka cluster context.

::

        {{.CLIName}} connector resume <id>
        {{.CLIName}} connector resume <id> --cluster <cluster-id>	`, c.Config.CLIName),
		RunE: c.resume,
		Args: cobra.ExactArgs(1),
	}
	cmd.Flags().String("cluster", "", "Kafka cluster ID.")
	cmd.Flags().SortFlags = false
	c.AddCommand(cmd)
}

func (c *command) list(cmd *cobra.Command, args []string) error {
	kafkaCluster, err := pcmd.KafkaCluster(cmd, c.Context)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	connectors, err := c.Client.Connect.ListWithExpansions(context.Background(), &connectv1.Connector{AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id}, "status,info,id")
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	outputWriter, err := output.NewListOutputWriter(cmd, listFields, listFields, listStructuredLabels)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	for name, connector := range connectors {
		connector := &describeDisplay{
			Name:   name,
			ID:     connector.Id.Id,
			Status: connector.Status.Connector.State,
			Type:   connector.Info.Type,
		}
		outputWriter.AddElement(connector)
	}
	return outputWriter.Out()
}

func (c *command) describe(cmd *cobra.Command, args []string) error {
	kafkaCluster, err := pcmd.KafkaCluster(cmd, c.Context)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	connector, err := c.Client.Connect.GetExpansionById(context.Background(), &connectv1.Connector{AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Id: args[0]})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	pcmd.Println(cmd, "Connector Details")
	data := &describeDisplay{
		Name:   connector.Status.Name,
		ID:     connector.Id.Id,
		Status: connector.Status.Connector.State,
		Type:   connector.Info.Type,
	}
	_ = printer.RenderTableOut(data, listFields, describeRenames, os.Stdout)

	pcmd.Println(cmd, "\n\nTask Level Details")
	var tasks [][]string
	titleRow := []string{"Task_ID", "State"}
	for _, task := range connector.Status.Tasks {
		record := &struct {
			Task_ID int32
			State   string
		}{
			task.Id,
			task.State,
		}
		tasks = append(tasks, printer.ToRow(record, titleRow))
	}
	printer.RenderCollectionTable(tasks, titleRow)
	pcmd.Println(cmd, "\n\nConfiguration Details")
	var configs [][]string
	titleRow = []string{"Configuration", "Value"}
	for name, value := range connector.Info.Config {
		record := &struct {
			Configuration string
			Value         string
		}{
			name,
			value,
		}
		configs = append(configs, printer.ToRow(record, titleRow))
	}
	printer.RenderCollectionTable(configs, titleRow)
	return nil
}

func (c *command) create(cmd *cobra.Command, args []string) error {
	kafkaCluster, err := pcmd.KafkaCluster(cmd, c.Context)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	userConfigs, err := getConfig(cmd)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	connector, err := c.Client.Connect.Create(context.Background(), &connectv1.ConnectorConfig{UserConfigs: *userConfigs, AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Name: (*userConfigs)["name"], Plugin: (*userConfigs)["connector.class"]})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	pcmd.Printf(cmd, "Created connector %s ", connector.Name)
	// Resolve Connector ID from Name of created connector
	connectorID, err := c.Client.Connect.GetExpansionByName(context.Background(), &connectv1.Connector{AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Name: connector.Name})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	pcmd.Println(cmd, connectorID.Id.Id)
	return nil
}

func (c *command) update(cmd *cobra.Command, args []string) error {
	userConfigs, err := getConfig(cmd)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	kafkaCluster, err := pcmd.KafkaCluster(cmd, c.Context)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	// Resolve Connector Name from ID
	connector, err := c.Client.Connect.GetExpansionById(context.Background(), &connectv1.Connector{AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Id: args[0]})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	_, err = c.Client.Connect.Update(context.Background(), &connectv1.ConnectorConfig{UserConfigs: *userConfigs, AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Name: connector.Info.Name, Plugin: (*userConfigs)["connector.class"]})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	pcmd.Println(cmd, "Updated connector "+args[0])
	return nil
}

func (c *command) delete(cmd *cobra.Command, args []string) error {
	kafkaCluster, err := pcmd.KafkaCluster(cmd, c.Context)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	connector, err := c.Client.Connect.GetExpansionById(context.Background(), &connectv1.Connector{AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Id: args[0]})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	err = c.Client.Connect.Delete(context.Background(), &connectv1.Connector{Name: connector.Info.Name, AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	pcmd.Println(cmd, "Successfully deleted connector")
	return nil
}

func (c *command) pause(cmd *cobra.Command, args []string) error {
	kafkaCluster, err := pcmd.KafkaCluster(cmd, c.Context)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	connector, err := c.Client.Connect.GetExpansionById(context.Background(), &connectv1.Connector{AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Id: args[0]})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	err = c.Client.Connect.Pause(context.Background(), &connectv1.Connector{Name: connector.Info.Name, AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	pcmd.Println(cmd, "Successfully paused connector")
	return nil
}

func (c *command) resume(cmd *cobra.Command, args []string) error {
	kafkaCluster, err := pcmd.KafkaCluster(cmd, c.Context)
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	connector, err := c.Client.Connect.GetExpansionById(context.Background(), &connectv1.Connector{AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id, Id: args[0]})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	err = c.Client.Connect.Resume(context.Background(), &connectv1.Connector{Name: connector.Info.Name, AccountId: c.EnvironmentId(), KafkaClusterId: kafkaCluster.Id})
	if err != nil {
		return errors.HandleCommon(err, cmd)
	}
	pcmd.Println(cmd, "Successfully resumed connector")
	return nil
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
