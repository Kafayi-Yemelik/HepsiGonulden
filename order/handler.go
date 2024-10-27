package order

import (
	"HepsiGonulden/order/types"
	"HepsiGonulden/validation"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type OrderHandler struct {
	service *Service
}

// @title Order Service API
// @version 1.0
// @description This is the Order Service API for handling CRUD operations related to order
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /
// @schemes http
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func NewHandler(f *fiber.App, service *Service) {
	handler := &OrderHandler{service: service}

	api := f.Group("/orders")

	api.Get("/:id", handler.GetByID)
	api.Post("/", handler.OrderCreate)
	api.Put("/:id", handler.Update)
	api.Delete("/:id", handler.Delete)
}

// GetByID retrieves an order by its ID.
// @Summary Get order by ID
// @Description Get order details by ID
// @Tags order
// @Produce  json
// @Param id path string true "Order ID"
// @Param  Authorization header string true "JWT token"
// @Success 200 {object} types.OrderResponseModel
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [get]
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
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	creatorUserId := claims["Id"].(string)
	orderRequestModel.CreatorUserId = creatorUserId

	id, err := h.service.CreateOrder(c.Context(), &orderRequestModel)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"orderId": id,
		"message": "Order created successfully",
	})

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
