package cmd

import (
	"HepsiGonulden/Customer"
	"HepsiGonulden/mongo"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"time"
)

func CustomerApiCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "customer",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fiber.New()

			mongoClient, err := mongo.GetMongoClient(10 * time.Second)
			if err != nil {
				return err
			}
			repo, err := Customer.NewRepository(mongoClient)
			if err != nil {
				return err
			}
			service := Customer.NewService(repo)

			/*
				1. Mongo client oluşturulması
				2. Repository oluşturulması
				3. Service oluşturulması
				4. Servicelerin handlera verilmesi
				5. Customer endpointlerinin içinin yazılması
			*/
			Customer.NewHandler(app, service)

			app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!, Customer")
			})

			app.Listen(":3000")
			return nil
		},
	}

	return rootCmd
}
