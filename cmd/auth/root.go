package auth

import "github.com/spf13/cobra"

func NewAuthCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "auth",
	}
	command.AddCommand(AuthApiCommand())
	return command
}
