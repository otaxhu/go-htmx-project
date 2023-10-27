package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/repository"
	"github.com/otaxhu/go-htmx-project/internal/service"
	"github.com/otaxhu/go-htmx-project/internal/web"
)

//go:embed .env
var envVarsFile []byte

//go:embed static/*
var staticFiles embed.FS

func main() {
	closeAppSignal := make(chan os.Signal, 1)
	signal.Notify(closeAppSignal, os.Interrupt, syscall.SIGTERM)

	// Closeables dependencies
	var (
		productsRepo repository.ProductsRepository
		webApp       web.WebApp
	)

	ctx := context.Background()

	// Gracefully shutting down app when closeAppSignal gets a signal
	go func() {
		<-closeAppSignal
		if webApp != nil {
			fmt.Println("Apagando WebApp...")
			if err := webApp.Shutdown(ctx); err != nil {
				log.Fatal(err)
			}
		}

		if productsRepo != nil {
			fmt.Println("Apagando ProductsRepository...")
			if err := productsRepo.Close(); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Aplicacion apagada exitosamente")
		os.Exit(0)
	}()

	// Config DI
	cfg, err := config.New(envVarsFile)
	if err != nil {
		log.Fatal(err)
	}

	// Repositories DI
	productsRepo, err = repository.NewProductsRepository(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	_ = repository.NewImageRepository(cfg.ImageRepo)

	// Services DI
	productsService := service.NewProductsService(cfg.ProductsService, productsRepo)

	// Web Framework DI
	webApp, err = web.NewWebApp(cfg.Server, productsService, staticFiles)
	if err != nil {
		log.Fatal(err)
	}

	webAppStartedSignal := make(chan struct{})
	go func() {
		if err := webApp.StartAndNotify(webAppStartedSignal); err != nil {
			log.Fatal(err)
		}
	}()
	<-webAppStartedSignal
	fmt.Println("Aplicacion inicializada exitosamente")
	var blockingCh chan struct{}
	<-blockingCh
}
