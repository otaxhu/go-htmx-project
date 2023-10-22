package chi

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *chiApp) bindRoutes() {
	r := chi.NewRouter()

	// Global middlewares
	r.Use(middleware.RedirectSlashes)

	r.Route("/api/v1/products", func(r chi.Router) {

		// Group of routes that returns application/json as response
		r.Group(func(r chi.Router) {
			r.Use(middleware.SetHeader("Content-Type", "application/json"))

			r.Get("/", app.productsHandlers.GetAndSearchProducts)
			r.Get("/{id}", app.productsHandlers.GetProductById)
			r.Post("/", app.productsHandlers.PostProduct)
			r.Put("/{id}", app.productsHandlers.PutProduct)
		})

		r.Delete("/{id}", app.productsHandlers.DeleteProductById)
	})
	app.server.Handler = r
}
