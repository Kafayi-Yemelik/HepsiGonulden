package Customer

import "github.com/gofiber/fiber/v2"

type CustomerHandler struct{}

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
	return c.SendString("Get")
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
