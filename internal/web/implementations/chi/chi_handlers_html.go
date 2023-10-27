package chi

import (
	"fmt"
	"net/http"

	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/templates/pages"
)

type chiHtmlProductsHandlers struct {
	productsService service.ProductsService
}

func (handler *chiHtmlProductsHandlers) GetAndSearchProducts(w http.ResponseWriter, r *http.Request) {
	
	if err := pages.HomeProducts("Hola Mundo!!!").Render(r.Context(), w); err != nil {
		fmt.Fprintf(w, "there was an error trying to rendering HomeProducts: %v", err)
	}
}

func (handler *chiHtmlProductsHandlers) GetProductById(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (handler *chiHtmlProductsHandlers) PostProduct(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (handler *chiHtmlProductsHandlers) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (handler *chiHtmlProductsHandlers) PutProduct(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}
