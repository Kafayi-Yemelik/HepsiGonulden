package order

import (
	"HepsiGonulden/order/types"
	"HepsiGonulden/validation"
	"context"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service *Service
}

func NewHandler(f *fiber.App, service *Service) {
	handler := &OrderHandler{service: service}

	api := f.Group("/orders")

	api.Get("/:id", handler.GetByID)
	api.Post("/", handler.OrderCreate)
	api.Put("/:id", handler.Update)
	api.Delete("/:id", handler.Delete)
}
func (h *OrderHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	order, err := h.service.GetById(context.Background(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	orderResponse := types.ToOrderResponse(order)

	return c.Status(fiber.StatusOK).JSON(orderResponse)
}

func (h *OrderHandler) OrderCreate(c *fiber.Ctx) error {

	var orderRequestModel types.OrderRequestModel
	if err := c.BodyParser(&orderRequestModel); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if err := validation.Validate(orderRequestModel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// ToDo Gönül : CreaterUser IdYe nasıl ulaşacağız
	id, err := h.service.CreateOrder(c.Context(), &orderRequestModel)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}
func (h *OrderHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var order types.OrderUpdateModel

	if err := c.BodyParser(&order); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else if err := validation.Validate(order); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.service.Update(c.Context(), id, order); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(order)

}

func (h *OrderHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
