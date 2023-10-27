package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *chiApp) bindRoutesAndHandlers() {
	r := chi.NewRouter()
	// Global middlewares
	r.Use(middleware.RedirectSlashes)
	r.Get("/api/static/*", http.StripPrefix("/api", http.FileServer(http.FS(app.staticFS))).ServeHTTP)

	r.Route("/products", func(r chi.Router) {

		r.Get("/", app.productsHandlers.GetAndSearchProducts)
		r.Get("/{id}", app.productsHandlers.GetProductById)
		r.Post("/", app.productsHandlers.PostProduct)
		r.Put("/{id}", app.productsHandlers.PutProduct)
		r.Delete("/{id}", app.productsHandlers.DeleteProductById)

	})
	app.server.Handler = r
}
