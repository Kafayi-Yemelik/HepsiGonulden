package api

import (
	"HepsiGonulden/internal/handler"
	"HepsiGonulden/internal/repository"
	"HepsiGonulden/internal/services"
	"HepsiGonulden/pkg/mongo"
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
			repo, err := repository.NewCustomerRepository(mongoClient)
			if err != nil {
				return err
			}

			service := services.NewCustomerService(repo)
			handler.NewAuthHandler(app, service)

			app.Listen(":3002")
			return nil
		},
	}
	return rootCmd
}
