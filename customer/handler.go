package customer

import (
	"HepsiGonulden/customer/types"
	"HepsiGonulden/validation"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CustomerHandler struct {
	service *Service
}

// @title Customer Service API
// @version 1.0
// @description This is the Customer Service API for handling CRUD operations related to customer
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

func NewHandler(f *fiber.App, service *Service) {
	handler := &CustomerHandler{service: service}

	api := f.Group("/customers")

	api.Get("/:id", handler.GetByID)
	api.Post("/", handler.Create)
	api.Put("/:id", handler.Update)
	api.Delete("/:id", handler.Delete)
}

// GetByID retrieves an customer by its ID.
// @Summary Get customer by ID
// @Description Get customer details by ID
// @Tags customer
// @Produce  json
// @Param id path string true "Customer ID"
// @Param  Authorization header string true "JWT token"
// @Success 200 {object} types.CustomerResponseModel
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetByID(c *fiber.Ctx) error {

	id := c.Params("id")

	customer, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	customerResponse := types.ToCustomerResponse(customer)

	return c.Status(fiber.StatusOK).JSON(customerResponse)

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
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	creatorUserId := claims["Id"].(string)
	customerRequestModel.CreatorUserId = creatorUserId

	customerID, err := h.service.Create(c.Context(), customerRequestModel)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(customerID)
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
