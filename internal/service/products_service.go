package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/models/dto"
	"github.com/otaxhu/go-htmx-project/internal/repository"
)

type ProductsService interface {
	// errors that this function can return:
	//
	// 1. ErrNotFound
	//
	// 2. ErrInternalServer
	//
	// 3. ErrInvalidInput
	GetProducts(ctx context.Context, page int) ([]dto.GetProduct, error)

	// errors that this function can return:
	//
	// 1. ErrInternalServer
	//
	// 2. ErrInvalidInput
	//
	// 3. ErrNotFound
	GetProductById(ctx context.Context, id string) (dto.GetProduct, error)

	// errors that this function can return:
	//
	// 1. ErrInternalServer
	//
	// 2. ErrInvalidInput
	SaveProduct(ctx context.Context, product dto.SaveProduct) (dto.GetProduct, error)

	// errors that this function can return:
	//
	// 1. ErrInternalServer
	//
	// 2. ErrInvalidInput
	//
	// 3. ErrNotFound
	SearchProducts(ctx context.Context, term string) ([]dto.GetProduct, error)

	// errors that this function can return:
	//
	// 1. ErrInternalServer
	//
	// 2. ErrInvalidInput
	//
	// 3. ErrNotFound
	UpdateProduct(ctx context.Context, product dto.UpdateProduct) (dto.GetProduct, error)

	// errors that this function can return:
	//
	// 1. ErrInternalServer
	//
	// 2. ErrInvalidInput
	//
	// 3. ErrNotFound
	DeleteProductById(ctx context.Context, id string) error
}

func NewProductsService(productsServiceCfg config.ProductsService, productsRepo repository.ProductsRepository) ProductsService {
	return &productsServiceImpl{
		cfg:          productsServiceCfg,
		productsRepo: productsRepo,
		validate:     validator.New(),
	}
}
