package chi

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"

	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/interfaces"
	"github.com/otaxhu/go-htmx-project/settings"
)

type chiApp struct {
	server           *http.Server
	viewsFS          embed.FS
	templateFuncs    template.FuncMap
	productsService  service.ProductsService
	productsHandlers interfaces.ProductsHandlers
}

func NewChiApp(serverSettings settings.Server, productsService service.ProductsService) *chiApp {
	return &chiApp{
		server: &http.Server{
			Addr: ":" + strconv.Itoa(int(serverSettings.Port)),
		},
		productsService: productsService,
	}
}

func (app *chiApp) Start() error {
	var err error
	app.productsHandlers, err = newChiProductsHandlers(app.productsService, app.viewsFS, app.templateFuncs)
	if err != nil {
		return err
	}
	app.bindRoutes()
	return app.server.ListenAndServe()
}

func (app *chiApp) SetViews(viewsFS embed.FS) {
	app.viewsFS = viewsFS
}

func (app *chiApp) SetTemplateFuncs(funcs template.FuncMap) {
	app.templateFuncs = funcs
}
