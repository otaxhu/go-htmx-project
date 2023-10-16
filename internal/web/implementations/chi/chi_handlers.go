package chi

import (
	"github.com/otaxhu/go-htmx-project/internal/service"
)

type chiProductsHandlers struct {
	productsService service.ProductsService
}

func newChiProductsHandlers(productsService service.ProductsService) *chiProductsHandlers {
	return &chiProductsHandlers{
		productsService: productsService,
	}
}
