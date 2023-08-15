package service

import "github.com/otaxhu/go-htmx-project/internal/repository"

type ProductsService interface{}

type productsServiceImpl struct {
	productsRepo repository.ProductsRepository
}

func NewProductsService(productsRepo repository.ProductsRepository) ProductsService {
	return &productsServiceImpl{
		productsRepo: productsRepo,
	}
}
