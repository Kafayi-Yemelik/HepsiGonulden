package cmd

import (
	"HepsiGonulden/customer"
	"HepsiGonulden/mongo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/cobra"
	"time"
)

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func CustomerApiCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "customer",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fiber.New(fiber.Config{
				// Global custom error handler
				ErrorHandler: func(c *fiber.Ctx, err error) error {
					return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
						Success: false,
						Message: err.Error(),
					})
				},
			})
			app.Use(requestid.New())
			app.Use(logger.New(logger.Config{
				Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
			}))
			mongoClient, err := mongo.GetMongoClient(10 * time.Second)
			if err != nil {
				return err
			}
			repo, err := customer.NewRepository(mongoClient)
			if err != nil {
				return err
			}
			service := customer.NewService(repo)

			/*
				1. Mongo client oluşturulması
				2. Repository oluşturulması
				3. Service oluşturulması
				4. Servicelerin handlera verilmesi
				5. customer endpointlerinin içinin yazılması
			*/
			customer.NewHandler(app, service)

			app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!, customer")
			})

			app.Listen(":3000")
			return nil
		},
	}

	return rootCmd
}
