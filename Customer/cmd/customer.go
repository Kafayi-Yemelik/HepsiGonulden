package cmd

import (
	"HepsiGonulden/Customer"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CustomerApiCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "customer",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fiber.New()

			/*
				1. Mongo client oluşturulması
				2. Repository oluşturulması
				3. Service oluşturulması
				4. Servicelerin handlera verilmesi
				5. Customer endpointlerinin içinin yazılması
			*/
			Customer.NewHandler(app)

			app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!, Customer")
			})

			app.Listen(":3000")
			return nil
		},
	}

	return rootCmd
}
