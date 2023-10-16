package web

import (
	"fmt"

	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/service"
	chi_implementation "github.com/otaxhu/go-htmx-project/internal/web/implementations/chi"
)

type WebApp interface {
	Start() error
}

func NewWebApp(serverCfg config.Server, productsService service.ProductsService) (WebApp, error) {
	switch serverCfg.Framework {
	case "chi":
		return chi_implementation.NewChiApp(serverCfg, productsService), nil
	default:
		return nil, fmt.Errorf("the `%s` framework does not have a `WebApp` implementation", serverCfg.Framework)
	}
}
