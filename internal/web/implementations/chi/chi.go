package chi

import (
	"fmt"
	"net/http"

	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/interfaces"
)

type chiApp struct {
	server           *http.Server
	productsService  service.ProductsService
	productsHandlers interfaces.ProductsHandlers
}

func NewChiApp(serverCfg config.Server, productsService service.ProductsService) *chiApp {
	return &chiApp{
		server: &http.Server{
			Addr: ":" + fmt.Sprint(serverCfg.Port),
		},
		productsService: productsService,
	}
}

func (app *chiApp) Start() error {
	app.productsHandlers = newChiProductsHandlers(app.productsService)
	app.bindRoutes()
	return app.server.ListenAndServe()
}
