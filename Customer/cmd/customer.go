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
