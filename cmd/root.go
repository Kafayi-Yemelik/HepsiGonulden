package cmd

import (
	"HepsiGonulden/cmd/api"
	"HepsiGonulden/cmd/consumer"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Short: "",
	}
	rootCmd.AddCommand(api.CustomerApiCommand())
	rootCmd.AddCommand(api.OrderApiCommand())
	rootCmd.AddCommand(api.AuthApiCommand())
	// rootCmd.AddCommand(api.NewApiCommand())
	rootCmd.AddCommand(consumer.NewConsumerCommand())
	return rootCmd
}
