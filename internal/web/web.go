package web

import (
	"embed"
	"fmt"
	"html/template"

	"github.com/otaxhu/go-htmx-project/internal/service"
	chi_implementation "github.com/otaxhu/go-htmx-project/internal/web/implementations/chi"
	"github.com/otaxhu/go-htmx-project/settings"
)

type WebApp interface {
	SetViews(viewsFS embed.FS)
	SetTemplateFuncs(funcs template.FuncMap) error
	Start() error
}

func NewWebApp(serverSettings settings.Server, productsService service.ProductsService) (WebApp, error) {
	switch serverSettings.Framework {
	case "chi":
		return chi_implementation.NewChiApp(serverSettings, productsService), nil
	default:
		return nil, fmt.Errorf("the `%s` framework does not have a `WebApp` implementation", serverSettings.Framework)
	}
}
