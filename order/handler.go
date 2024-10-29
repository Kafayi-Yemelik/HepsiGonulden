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

// OrderCreate creates a new order.
// @Summary Create a new order
// @Description Create a new order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body types.OrderRequestModel true "Order data"
// @Param Authorization header string true "JWT token"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders [post]
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

// Update modifies an existing order's details.
// @Summary Update order details
// @Description Update order details with the given data
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Param Authorization header string true "JWT token"
// @Param order body types.OrderUpdateModel true "Order data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /orders/{id} [put]
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

// Delete removes an order from the database.
// @Summary Delete order
// @Description Delete an order by its ID
// @Tags order
// @Produce  json
// @Param Authorization header string true "JWT token"
// @Param id path string true "Order ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} string
// @Router /orders/{id} [delete]
func (h *OrderHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(c.Context(), id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
