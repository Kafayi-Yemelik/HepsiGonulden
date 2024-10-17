package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func OrderApiCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "order",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fiber.New()

			app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!, order")
			})

			app.Listen(":3001")
			return nil
		},
	}

	return rootCmd
}
