package chi

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"

	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/std_http_interfaces"
	"github.com/otaxhu/go-htmx-project/settings"
)

type chiApp struct {
	server   *http.Server
	viewsFS  embed.FS
	handlers std_http_interfaces.Handlers
}

func NewChiApp(serverSettings settings.Server, productsService service.ProductsService) *chiApp {
	return &chiApp{
		server: &http.Server{
			Addr: ":" + strconv.Itoa(int(serverSettings.Port)),
		},
	}
}

func (app *chiApp) Start() error {
	app.bindRoutes()
	return app.server.ListenAndServe()
}

func (app *chiApp) SetViews(viewsFS embed.FS) {
	app.viewsFS = viewsFS
}

func (app *chiApp) SetTemplateFuncs(funcs template.FuncMap) error {
	var err error
	app.handlers, err = newChiHandlers(app.viewsFS, funcs)
	if err != nil {
		return err
	}
	return nil
}
