package main

import (
	"log"

	"github.com/otaxhu/go-htmx-project/internal/repository"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web"
	"github.com/otaxhu/go-htmx-project/settings"
)

func main() {
	serverSettings, err := settings.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	dbSettings, err := settings.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	productsRepo, err := repository.NewProductsRepository(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	imageRepo := repository.NewImageRepository()
	productsService := service.NewProductsService(productsRepo, imageRepo)
	app, err := web.NewWebApp(serverSettings, productsService)
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// Go generate for generating TailwindCSS output file
//
//go:generate ./tailwind.exe -i ./tailwind.input.css -o ./public/css/tailwind.css
