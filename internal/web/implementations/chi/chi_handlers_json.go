package chi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/elnormous/contenttype"
	"github.com/go-chi/chi/v5"
	"github.com/otaxhu/go-htmx-project/internal/models/dto"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web/helpers"
)

type chiJsonProductsHandlers struct {
	productsService service.ProductsService
}

func (handler *chiJsonProductsHandlers) GetAndSearchProducts(w http.ResponseWriter, r *http.Request) {
	// SearchProducts
	products, err := handler.productsService.SearchProducts(r.Context(), r.URL.Query().Get("searchTerm"))
	if err == service.ErrNotFound {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	} else if err == service.ErrInternalServer {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	} else if err == nil {
		if err := json.NewEncoder(w).Encode(products); err != nil {
			helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	// GetProducts only if searchTerm query param is not present
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}
	products, err = handler.productsService.GetProducts(r.Context(), page)
	if err == service.ErrNotFound {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	} else if err == service.ErrInvalidInput {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(products); err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *chiJsonProductsHandlers) GetProductById(w http.ResponseWriter, r *http.Request) {
	product, err := handler.productsService.GetProductById(r.Context(), chi.URLParam(r, "id"))
	if err == service.ErrInvalidInput {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	} else if err == service.ErrNotFound {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(product); err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *chiJsonProductsHandlers) PostProduct(w http.ResponseWriter, r *http.Request) {
	p := dto.SaveProduct{}
	mediaType, err := contenttype.GetMediaType(r)
	if err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType.EqualsMIME(mtJson) {
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else if mediaType.EqualsMIME(mtForm) || mediaType.EqualsMIME(mtMultiForm) {
		p.Name = r.FormValue("name")
		p.Description = r.FormValue("description")
	} else {
		helpers.JsonErrorResponse(w, unsupportedMediaTypeErr, http.StatusUnsupportedMediaType)
		return
	}
	createdProduct, err := handler.productsService.SaveProduct(r.Context(), p)
	if err == service.ErrInvalidInput {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdProduct); err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *chiJsonProductsHandlers) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	if err := handler.productsService.DeleteProductById(r.Context(), chi.URLParam(r, "id")); err == service.ErrInvalidInput {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	} else if err == service.ErrNotFound {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Del("Content-Type")
	w.WriteHeader(http.StatusNoContent)
}

func (handler *chiJsonProductsHandlers) PutProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	p := dto.UpdateProduct{Id: id}
	mediaType, err := contenttype.GetMediaType(r)
	if err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType.EqualsMIME(mtJson) {
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else if mediaType.EqualsMIME(mtForm) || mediaType.EqualsMIME(mtMultiForm) {
		p.Name = r.FormValue("name")
		p.Description = r.FormValue("description")
	} else {
		helpers.JsonErrorResponse(w, unsupportedMediaTypeErr, http.StatusUnsupportedMediaType)
		return
	}
	updatedProduct, err := handler.productsService.UpdateProduct(r.Context(), p)
	if err == service.ErrNotFound {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	} else if err == service.ErrInvalidInput {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(updatedProduct); err != nil {
		helpers.JsonErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
