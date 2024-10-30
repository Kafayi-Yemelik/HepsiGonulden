package cmd

import (
	"HepsiGonulden/kafka"
	"HepsiGonulden/mongo"
	"HepsiGonulden/order"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/cobra"
	"time"
)

type GlobalErrorHandlerResp struct {
	Success bool
	Message string
}

func OrderApiCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "order",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fiber.New(fiber.Config{
				// Global custom error handler
				ErrorHandler: func(c *fiber.Ctx, err error) error {
					fiberErr, ok := err.(*fiber.Error)
					if ok {
						return c.Status(fiberErr.Code).JSON(GlobalErrorHandlerResp{
							Success: false,
							Message: fiberErr.Message,
						})
					}

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
			cfg := swagger.Config{
				BasePath: "/",
				FilePath: "./order/docs/swagger.json",
				Path:     "swagger",
				Title:    "Swagger API Docs",
			}

			app.Use(swagger.New(cfg))
			app.Use(jwtware.New(jwtware.Config{
				SigningKey: jwtware.SigningKey{Key: []byte("secret")},
			}))
			mongoClient, err := mongo.GetMongoClient(10 * time.Second)
			if err != nil {
				return err
			}
			repo, err := order.NewRepository(mongoClient)
			if err != nil {
				return err
			}

			producer := kafka.NewProducer()
			service := order.NewService(repo, producer)

			order.NewHandler(app, service)
			app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!, order")
			})

			app.Listen(":3001")
			return nil
		},
	}

	return rootCmd
}
