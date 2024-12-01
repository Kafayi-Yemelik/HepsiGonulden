package cmd

import (
	"HepsiGonulden/cmd/auth"
	"HepsiGonulden/cmd/consumer"
	"HepsiGonulden/cmd/customer"
	"HepsiGonulden/cmd/order"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Short: "",
	}
	rootCmd.AddCommand(consumer.NewConsumerCommand())
	rootCmd.AddCommand(order.NewOrderCommand())
	rootCmd.AddCommand(auth.NewAuthCommand())
	rootCmd.AddCommand(customer.NewCustomerCommand())
	return rootCmd
}
