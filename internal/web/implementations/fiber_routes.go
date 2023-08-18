package implementations

func (a *fiberApp) bindRoutes() {
	a.app.Post("/submit", a.productsHandlers.Submit)
	a.app.Get("/", a.productsHandlers.Home)
	a.app.Static("/static", "./public")
}
