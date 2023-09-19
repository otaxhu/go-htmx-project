package chi

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/otaxhu/go-htmx-project/internal/models/dto"
)

type chiHandlers struct {
	*template.Template
}

func newChiHandlers(viewsFS embed.FS, funcs template.FuncMap) (*chiHandlers, error) {
	tmpls, err := template.New("").Funcs(funcs).ParseFS(viewsFS, "internal/web/views/*/*.html", "internal/web/views/*/*/*.html")
	if err != nil {
		return nil, err
	}
	return &chiHandlers{
		Template: tmpls,
	}, nil
}

func (handler *chiHandlers) HomeGetProducts(w http.ResponseWriter, r *http.Request) {
	products := []dto.GetProduct{}

	for i := 1; i <= 4; i++ {
		products = append(products, dto.GetProduct{Name: "Producto " + fmt.Sprint(i), Description: "Descripcion " + fmt.Sprint(i)})
	}

	w.Header().Set("Content-Type", "text/html")

	if r.Header.Get("HX-Request") != "true" {
		handler.ExecuteTemplate(w, "pages/home", map[string]any{"products": products})
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page <= 0 {
		handler.ExecuteTemplate(w, "pages/home", map[string]any{"products": products})
		return
	}

	handler.ExecuteTemplate(w, "pages/hx/home", map[string]any{"products": products})
}
