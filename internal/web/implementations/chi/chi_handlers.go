package chi

import (
	"net/http"

	"github.com/elnormous/contenttype"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/handlers"
)

type chiProductsHandlers struct {
	jsonHandlers handlers.ProductsHandlers
	htmlHandlers handlers.ProductsHandlers
}

func newChiProductsHandlers(productsService service.ProductsService) *chiProductsHandlers {
	return &chiProductsHandlers{
		jsonHandlers: &chiJsonProductsHandlers{
			productsService: productsService,
		},
		htmlHandlers: &chiHtmlProductsHandlers{
			productsService: productsService,
		},
	}
}

// MediaTypes
var (
	mtHtml = contenttype.NewMediaType("text/html")
	mtJson = contenttype.NewMediaType("application/json")

	// These are the available representations of this endpoint [/products]
	availableMediaTypes = []contenttype.MediaType{
		mtHtml,
		mtJson,
	}
)

func (handler *chiProductsHandlers) GetAndSearchProducts(w http.ResponseWriter, r *http.Request) {
	accepted, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if err == contenttype.ErrNoAcceptableTypeFound {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	setHeaderMw := middleware.SetHeader("Content-Type", accepted.MIME())
	if accepted.EqualsMIME(mtHtml) {
		setHeaderMw(http.HandlerFunc(handler.htmlHandlers.GetAndSearchProducts)).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		setHeaderMw(http.HandlerFunc(handler.jsonHandlers.GetAndSearchProducts)).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) GetProductById(w http.ResponseWriter, r *http.Request) {
	accepted, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if err == contenttype.ErrNoAcceptableTypeFound {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	setHeaderMw := middleware.SetHeader("Content-Type", accepted.MIME())
	if accepted.EqualsMIME(mtHtml) {
		setHeaderMw(http.HandlerFunc(handler.htmlHandlers.GetProductById)).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		setHeaderMw(http.HandlerFunc(handler.jsonHandlers.GetProductById)).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) PostProduct(w http.ResponseWriter, r *http.Request) {
	accepted, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if err == contenttype.ErrNoAcceptableTypeFound {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	setHeaderMw := middleware.SetHeader("Content-Type", accepted.MIME())
	if accepted.EqualsMIME(mtHtml) {
		setHeaderMw(http.HandlerFunc(handler.htmlHandlers.PostProduct)).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		setHeaderMw(http.HandlerFunc(handler.jsonHandlers.PostProduct)).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	accepted, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if err == contenttype.ErrNoAcceptableTypeFound {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	setHeaderMw := middleware.SetHeader("Content-Type", accepted.MIME())
	if accepted.EqualsMIME(mtHtml) {
		setHeaderMw(http.HandlerFunc(handler.htmlHandlers.DeleteProductById)).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		setHeaderMw(http.HandlerFunc(handler.jsonHandlers.DeleteProductById)).ServeHTTP(w, r)
		return
	}
}

func (handler *chiProductsHandlers) PutProduct(w http.ResponseWriter, r *http.Request) {
	accepted, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)
	if err == contenttype.ErrNoAcceptableTypeFound {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	setHeaderMw := middleware.SetHeader("Content-Type", accepted.MIME())
	if accepted.EqualsMIME(mtHtml) {
		setHeaderMw(http.HandlerFunc(handler.htmlHandlers.PutProduct)).ServeHTTP(w, r)
		return
	} else if accepted.EqualsMIME(mtJson) {
		setHeaderMw(http.HandlerFunc(handler.jsonHandlers.PutProduct)).ServeHTTP(w, r)
		return
	}
}
