package main

import (
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/repository"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web"
)

//go:embed .env
var envVarsFile []byte

func main() {
	// Config DI
	cfg, err := config.New(envVarsFile)
	if err != nil {
		log.Fatal(err)
	}

	// Repositories DI
	productsRepo, err := repository.NewProductsRepository(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	_ = repository.NewImageRepository(cfg.ImageRepo)

	// Services DI
	productsService := service.NewProductsService(cfg.ProductsService, productsRepo)

	// Web Framework DI
	app, err := web.NewWebApp(cfg.Server, productsService)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := app.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("app running in: http://127.0.0.1:%d\n", cfg.Server.Port)
	var c chan int
	<-c
}
