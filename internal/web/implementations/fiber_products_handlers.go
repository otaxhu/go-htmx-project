package implementations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaxhu/go-htmx-project/internal/service"
)

type fiberProductsHandlers struct {
	productsService service.ProductsService
}

func newFiberProductsHandlers(productsService service.ProductsService) *fiberProductsHandlers {
	return &fiberProductsHandlers{
		productsService: productsService,
	}
}

func (handler *fiberProductsHandlers) Submit(c *fiber.Ctx) error {
	name := c.FormValue("name")
	return c.Render("result", name)
}

func (handler *fiberProductsHandlers) Home(c *fiber.Ctx) error {
	return c.Render("index.html", nil)
}
