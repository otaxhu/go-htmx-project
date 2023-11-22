package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/otaxhu/go-htmx-project/internal/web/middlewares"
)

func (app *chiApp) bindRoutesAndHandlers() {
	r := chi.NewRouter()
	// Global middlewares
	r.Use(middleware.RedirectSlashes)

	r.Get("/api/static/*", http.StripPrefix("/api", http.FileServer(http.FS(app.staticFS))).ServeHTTP)

	// CRUD api of products, process application/json, application/x-www-form-urlencoded,
	// multipart/form-data and respond with application/json or text/html depending on the Accept header
	r.Route("/api/products", func(r chi.Router) {

		r.Use(middlewares.FilterAcceptHeaderAndSetContentType(availableMediaTypes))

		r.Get("/", app.productsHandlers.GetAndSearchProducts)
		r.Get("/{id}", app.productsHandlers.GetProductById)
		r.Post("/", app.productsHandlers.PostProduct)
		r.Put("/{id}", app.productsHandlers.PutProduct)
		r.Delete("/{id}", app.productsHandlers.DeleteProductById)

	})

	// Requesting for templates, not the actual api.
	//
	// Always return text/html independently of the Accept header, so there is no reason for implement
	// content negotiation if there is only one type of content
	r.Group(func(r chi.Router) {

		r.Use(middleware.SetHeader("Content-Type", "text/html"))

		r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
			if err := app.templates.ExecuteTemplate(w, "pages/home", nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})

		r.Get("/publish", func(w http.ResponseWriter, r *http.Request) {
			if err := app.templates.ExecuteTemplate(w, "pages/publish", nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	})

	app.server.Handler = r
}
