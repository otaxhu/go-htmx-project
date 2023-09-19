package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *chiApp) bindRoutes() {
	r := chi.NewRouter()
	r.Get("/", app.handlers.HomeGetProducts)
	r.Handle("/api/static/*", http.StripPrefix("/api/static", http.FileServer(http.Dir("public/"))))
	app.server.Handler = r
}
