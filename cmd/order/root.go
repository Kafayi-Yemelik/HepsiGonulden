package order

import "github.com/spf13/cobra"

func NewOrderCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "order",
	}
	command.AddCommand(OrderApiCommand())
	return command
}
