package chi

import (
	"embed"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/otaxhu/go-htmx-project/internal/models/dto"
	"github.com/otaxhu/go-htmx-project/internal/service"
)

type chiProductsHandlers struct {
	*template.Template
	productsService service.ProductsService
}

func newChiProductsHandlers(productsService service.ProductsService, viewsFS embed.FS, funcs template.FuncMap) (*chiProductsHandlers, error) {
	tmpls, err := template.New("").Funcs(funcs).ParseFS(viewsFS, "internal/web/views/*/*.html")
	if err != nil {
		return nil, err
	}
	return &chiProductsHandlers{
		Template:        tmpls,
		productsService: productsService,
	}, nil
}

func (handler *chiProductsHandlers) GetHomeProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	var requestingMoreProducts bool

	if page <= 0 {
		page = 1
	} else {
		requestingMoreProducts = true
	}

	data := make(map[string]any, 5)

	data["htmx"] = r.Header.Get("HX-Request")
	data["nextPage"] = page + 1
	data["requestingMoreProducts"] = requestingMoreProducts

	products, err := handler.productsService.GetProducts(r.Context(), uint(page))
	if err == service.ErrNotFound {
		if data["htmx"] == "true" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		data["noMoreProducts"] = true
		w.WriteHeader(http.StatusNotFound)
		handler.ExecuteTemplate(w, "actions/GetProducts", data)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data["products"] = products

	if _, err := handler.productsService.GetProducts(r.Context(), uint(page)+1); err == service.ErrNotFound {
		data["noMoreProducts"] = true
	}

	handler.ExecuteTemplate(w, "actions/GetProducts", data)
}

func (handler *chiProductsHandlers) PostPublishProduct(w http.ResponseWriter, r *http.Request) {
	_, file, _ := r.FormFile("image")

	if file != nil {
		file.Filename = uuid.NewString() + filepath.Ext(file.Filename)
		file.Header.Set("image_url", "http://"+r.Host+"/api/static/images/products/"+file.Filename)
	}

	product := dto.SaveProduct{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Image:       file,
	}

	data := make(map[string]any, 3)
	data["htmx"] = r.Header.Get("HX-Request")
	data["postingProduct"] = true

	if err := handler.productsService.SaveProduct(r.Context(), product); err == service.ErrInvalidProductObject {
		w.WriteHeader(http.StatusBadRequest)
		handler.ExecuteTemplate(w, "actions/PostProduct", data)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data["success"] = true

	handler.ExecuteTemplate(w, "actions/PostProduct", data)
}

func (handler *chiProductsHandlers) GetPublishProductTemplate(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"htmx": r.Header.Get("HX-Request"),
	}

	handler.ExecuteTemplate(w, "pages/publish-product", data)
}
