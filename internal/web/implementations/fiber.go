package implementations

import (
	"embed"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/settings"
)

type fiberApp struct {
	port             uint
	app              *fiber.App
	productsHandlers *fiberProductsHandlers
}

func NewFiberApp(serverSettings settings.Server, productsService service.ProductsService, viewsFS embed.FS) *fiberApp {
	app := fiber.New(fiber.Config{
		Views: html.NewFileSystem(http.FS(viewsFS), ".html"),
	})
	return &fiberApp{
		port:             serverSettings.Port,
		app:              app,
		productsHandlers: newFiberProductsHandlers(productsService),
	}
}

func (a *fiberApp) Start() error {
	a.bindRoutes()
	return a.app.Listen(":" + strconv.Itoa(int(a.port)))
}
