package Customer

import (
	"HepsiGonulden/Customer/types"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CustomerHandler struct {
	service *Service
}

func NewHandler(f *fiber.App) {
	handler := &CustomerHandler{}

	api := f.Group("/customer")

	api.Get("/:id", handler.GetByID)
	api.Post("/", handler.Create)
	api.Put("/:id", handler.Update)
	api.Patch("/:id", handler.PartialUpdate)
	api.Delete("/:id", handler.Delete)
}

func (h *CustomerHandler) Create(c *fiber.Ctx) error {
	return c.SendString("Create")
}

func (h *CustomerHandler) GetByID(c *fiber.Ctx) error {

	id := c.Params("id")

	customer, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	customerResponse := types.ToCustomerResponse(customer)

	return c.Status(fiber.StatusOK).JSON(customerResponse)

}

func (h *CustomerHandler) Update(c *fiber.Ctx) error {
	return c.SendString("Update")
}

func (h *CustomerHandler) PartialUpdate(c *fiber.Ctx) error {
	return c.SendString(" Update")
}

func (h *CustomerHandler) Delete(c *fiber.Ctx) error {
	return c.SendString("Delet")
}
