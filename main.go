package main

import (
	"context"
	"fmt"
	"log"

	"github.com/otaxhu/go-htmx-project/internal/models"
	"github.com/otaxhu/go-htmx-project/internal/repository"
	"github.com/otaxhu/go-htmx-project/settings"
)

func main() {
	dbSettings, err := settings.NewDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	serverSettings, err := settings.NewServer()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dbSettings, serverSettings)
	productsRepo, err := repository.NewProductsRepository(dbSettings)
	if err != nil {
		log.Fatalln(err)
	}
	products := []models.Product{
		{
			Name:        "producto 1",
			Description: "description 1",
			ImageUrl:    "url 1",
		},
		{
			Name:        "producto 2",
			Description: "description 2",
			ImageUrl:    "url 2",
		},
		{
			Name:        "producto 3",
			Description: "description 3",
			ImageUrl:    "url 3",
		},
		{
			Name:        "producto 4",
			Description: "description 4",
			ImageUrl:    "url 4",
		},
	}
	for _, product := range products {
		productsRepo.UpsertProduct(context.TODO(), product)
	}
}
