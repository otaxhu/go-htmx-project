package chi

import (
	"html/template"
	"net/http"

	"github.com/elnormous/contenttype"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/handlers"
)

type chiProductsHandlers struct {
	jsonHandlers handlers.ProductsHandlers
	htmlHandlers handlers.ProductsHandlers
}

func newChiProductsHandlers(productsService service.ProductsService, templates *template.Template) *chiProductsHandlers {
	return &chiProductsHandlers{
		jsonHandlers: &chiJsonProductsHandlers{
			productsService: productsService,
		},
		htmlHandlers: &chiHtmlProductsHandlers{
			productsService: productsService,
			templates:       templates,
		},
	}
}

func (handler *chiProductsHandlers) GetAndSearchProducts(w http.ResponseWriter, r *http.Request) {
	accepted, _, _ := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if accepted.EqualsMIME(mtHtml) {
		http.HandlerFunc(handler.htmlHandlers.GetAndSearchProducts).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		http.HandlerFunc(handler.jsonHandlers.GetAndSearchProducts).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) GetProductById(w http.ResponseWriter, r *http.Request) {
	accepted, _, _ := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if accepted.EqualsMIME(mtHtml) {
		http.HandlerFunc(handler.htmlHandlers.GetProductById).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		http.HandlerFunc(handler.jsonHandlers.GetProductById).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) PostProduct(w http.ResponseWriter, r *http.Request) {
	accepted, _, _ := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if accepted.EqualsMIME(mtHtml) {
		http.HandlerFunc(handler.htmlHandlers.PostProduct).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		http.HandlerFunc(handler.jsonHandlers.PostProduct).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	accepted, _, _ := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if accepted.EqualsMIME(mtHtml) {
		http.HandlerFunc(handler.htmlHandlers.DeleteProductById).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		http.HandlerFunc(handler.jsonHandlers.DeleteProductById).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) PutProduct(w http.ResponseWriter, r *http.Request) {
	accepted, _, _ := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if accepted.EqualsMIME(mtHtml) {
		http.HandlerFunc(handler.htmlHandlers.PutProduct).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		http.HandlerFunc(handler.jsonHandlers.PutProduct).ServeHTTP(w, r)
		return
	}
}
