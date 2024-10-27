package cmd

import (
	authCmd "HepsiGonulden/auth/cmd"
	customerCmd "HepsiGonulden/customer/cmd"
	orderCmd "HepsiGonulden/order/cmd"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{

		Short: "",
	}
	rootCmd.AddCommand(customerCmd.CustomerApiCommand())
	rootCmd.AddCommand(orderCmd.OrderApiCommand())
	rootCmd.AddCommand(authCmd.AuthApiCommand())

	return rootCmd
}
