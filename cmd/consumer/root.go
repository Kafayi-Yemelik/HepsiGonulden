package consumer

import "github.com/spf13/cobra"

func NewConsumerCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "consumer",
	}
	command.AddCommand(NewOrderCreateConsumerCommand())
	return command
}
