package chi

import (
	"html/template"
	"net/http"

	"github.com/otaxhu/go-htmx-project/internal/service"
)

type chiHtmlProductsHandlers struct {
	productsService service.ProductsService
	templates       *template.Template
}

func (handler *chiHtmlProductsHandlers) GetAndSearchProducts(w http.ResponseWriter, r *http.Request) {
	if err := handler.templates.ExecuteTemplate(w, "actions/get-products", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
