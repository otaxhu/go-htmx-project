package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/otaxhu/go-htmx-project/internal/web/middlewares"
)

func (app *chiApp) bindRoutes() {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.SetHtmlContent)
		r.Get("/", app.productsHandlers.GetHomeProducts)
		r.Post("/publish", app.productsHandlers.PostPublishProduct)
		r.Get("/publish", app.productsHandlers.GetPublishProductTemplate)
	})
	r.Method(http.MethodGet, "/api/static/*", http.StripPrefix("/api/static", http.FileServer(http.Dir("public/"))))
	app.server.Handler = r
}
