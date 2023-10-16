package chi

import (
	"github.com/go-chi/chi/v5"
)

func (app *chiApp) bindRoutes() {
	r := chi.NewRouter()
	app.server.Handler = r
}
