package customer

import "github.com/spf13/cobra"

func NewCustomerCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "customer",
	}
	command.AddCommand(CustomerApiCommand())
	return command
}
