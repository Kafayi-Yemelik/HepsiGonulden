package handler

import (
	"HepsiGonulden/internal/services"
	"HepsiGonulden/internal/types"
	"HepsiGonulden/pkg/authentication"
	"context"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *services.CustomerService
}

func NewAuthHandler(f *fiber.App, service *services.CustomerService) {
	handler := &AuthHandler{service}
	api := f.Group("/auth")
	api.Post("/login", handler.Login)
}

func (h AuthHandler) Login(ctx *fiber.Ctx) error {
	var loginRequestModel types.LoginRequestModel
	if err := ctx.BodyParser(&loginRequestModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	customer, err := h.service.GetByEmail(context.Background(), loginRequestModel.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if customer == nil {
		return fiber.NewError(fiber.StatusNotFound, "No customer found with email")
	} else if loginRequestModel.Password != customer.Password {
		return fiber.NewError(fiber.StatusUnauthorized, "Incorrect password")
	}

	token, err := authentication.JwtGenerator(customer.Id, customer.FirstName, customer.LastName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{"token": token})
}
