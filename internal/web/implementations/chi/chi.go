package chi

import (
	"context"
	"embed"
	"fmt"
	"net"
	"net/http"

	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/handlers"
)

type chiApp struct {
	server           *http.Server
	productsService  service.ProductsService
	productsHandlers handlers.ProductsHandlers
	staticFS         embed.FS

	port    int
	started bool
}

func NewChiApp(serverCfg config.Server, productsService service.ProductsService, staticFS embed.FS) *chiApp {
	app := &chiApp{
		server:          &http.Server{},
		productsService: productsService,
		port:            serverCfg.Port,
		staticFS:        staticFS,
	}
	return app
}

func (app *chiApp) StartAndNotify(notifyChan chan struct{}) error {
	l, err := net.Listen("tcp", ":"+fmt.Sprint(app.port))
	if err != nil {
		return err
	}
	app.productsHandlers = newChiProductsHandlers(app.productsService)
	app.bindRoutesAndHandlers()
	app.started = true
	close(notifyChan)
	// Here blocks the goroutine until the server is Shutdown
	if err := app.server.Serve(l); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (app *chiApp) Shutdown(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}
