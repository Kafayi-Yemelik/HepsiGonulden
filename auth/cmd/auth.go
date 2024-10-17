package cmd

import (
	"HepsiGonulden/auth"
	"HepsiGonulden/customer"
	"HepsiGonulden/mongo"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"time"
)

func AuthApiCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "auth",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fiber.New()
			mongoClient, err := mongo.GetMongoClient(10 * time.Second)
			if err != nil {
				return err
			}
			repo, err := customer.NewRepository(mongoClient)
			if err != nil {
				return err
			}
			service := customer.NewService(repo)
			auth.NewHandler(app, service)

			app.Listen(":3002")
			return nil

		},
	}
	return rootCmd
}
