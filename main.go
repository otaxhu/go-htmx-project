package main

import (
	"embed"
	"log"

	"github.com/componentize-go/componentize"
	"github.com/otaxhu/go-htmx-project/internal/repository"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web"
	"github.com/otaxhu/go-htmx-project/settings"
)

//go:embed internal/web/views/*
var viewsFS embed.FS

func main() {
	// Settings DI
	serverSettings, err := settings.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	dbSettings, err := settings.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Repositories DI
	productsRepo, err := repository.NewProductsRepository(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	imageRepo := repository.NewImageRepository()

	// Services DI
	productsService := service.NewProductsService(productsRepo, imageRepo)

	// Web Framework DI
	app, err := web.NewWebApp(serverSettings, productsService)
	if err != nil {
		log.Fatal(err)
	}

	app.SetViews(viewsFS)

	tmplFuncs := componentize.Default()

	if err := app.SetTemplateFuncs(tmplFuncs); err != nil {
		log.Fatal(err)
	}
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
