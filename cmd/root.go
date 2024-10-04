package cmd

import (
	customerCmd "HepsiGonulden/Customer/cmd"
	orderCmd "HepsiGonulden/Order/cmd"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{

		Short: "",
	}
	rootCmd.AddCommand(customerCmd.CustomerApiCommand())
	rootCmd.AddCommand(orderCmd.OrderApiCommand())
	return rootCmd
}
