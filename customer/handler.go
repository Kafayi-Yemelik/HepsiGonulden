package customer

import (
	"HepsiGonulden/customer/types"
	"HepsiGonulden/validation"
	"context"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	service *Service
}

func NewHandler(f *fiber.App, service *Service) {
	handler := &CustomerHandler{service: service}

	api := f.Group("/customers")

	api.Get("/:id", handler.GetByID)
	api.Post("/", handler.Create)
	api.Put("/:id", handler.Update)
	api.Delete("/:id", handler.Delete)
}

func (h *CustomerHandler) Create(c *fiber.Ctx) error {
	var customerRequestModel types.CustomerRequestModel
	if err := c.BodyParser(&customerRequestModel); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if err := validation.Validate(customerRequestModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	customer, err := h.service.GetByEmail(context.Background(), customerRequestModel.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if customer != nil {
		return fiber.NewError(fiber.StatusConflict, "customer with this email already exists")
	}

	customerID, err := h.service.Create(c.Context(), customerRequestModel)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(customerID)
}

func (h *CustomerHandler) GetByID(c *fiber.Ctx) error {

	id := c.Params("id")

	customer, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	customerResponse := types.ToCustomerResponse(customer)

	return c.Status(fiber.StatusOK).JSON(customerResponse)

}

func (h *CustomerHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer types.CustomerUpdateModel

	if err := c.BodyParser(&customer); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if err := validation.Validate(customer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.service.Update(c.Context(), id, customer); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(customer)

}

func (h *CustomerHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
