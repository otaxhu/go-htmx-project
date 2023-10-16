package service

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/models"
	"github.com/otaxhu/go-htmx-project/internal/models/dto"
	"github.com/otaxhu/go-htmx-project/internal/repository"
	repo_errors "github.com/otaxhu/go-htmx-project/internal/repository/errors"
)

type ProductsService interface {
	GetProducts(ctx context.Context, page int) ([]dto.GetProduct, error)
	GetProductById(ctx context.Context, id string) (dto.GetProduct, error)
	SaveProduct(ctx context.Context, product dto.SaveProduct) (string, error)
	SearchProducts(ctx context.Context, term string) ([]dto.SaveProduct, error)
	UpdateProduct(ctx context.Context, product dto.UpdateProduct) error
	DeleteProduct(ctx context.Context, id string) error
}

type productsServiceImpl struct {
	cfg          config.ProductsService
	productsRepo repository.ProductsRepository
	validate     *validator.Validate
}

func NewProductsService(productsServiceCfg config.ProductsService, productsRepo repository.ProductsRepository) ProductsService {
	return &productsServiceImpl{
		cfg:          productsServiceCfg,
		productsRepo: productsRepo,
		validate:     validator.New(),
	}
}

func (serv *productsServiceImpl) GetProducts(ctx context.Context, page int) ([]dto.GetProduct, error) {
	if page <= 0 {
		return nil, ErrInvalidPageParam
	}
	page--
	offset := serv.cfg.GetProductsLimit * page
	products, err := serv.productsRepo.GetProducts(ctx, offset, serv.cfg.GetProductsLimit)
	if err == repo_errors.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		log.Println(err)
		return nil, ErrInternalServer
	}
	dtoProduct := []dto.GetProduct{}
	for _, product := range products {
		dtoProduct = append(dtoProduct, dto.GetProduct{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
		})
	}
	return dtoProduct, nil
}

func (serv *productsServiceImpl) GetProductById(ctx context.Context, id string) (dto.GetProduct, error) {
	product, err := serv.productsRepo.GetProductById(ctx, id)
	if err == repo_errors.ErrNoRows {
		return dto.GetProduct{}, ErrNotFound
	} else if err != nil {
		log.Println(err)
		return dto.GetProduct{}, ErrInternalServer
	}
	return dto.GetProduct{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
	}, nil
}

func (serv *productsServiceImpl) SearchProducts(ctx context.Context, term string) ([]dto.SaveProduct, error) {
	dbProducts, err := serv.productsRepo.SearchProducts(ctx, term)
	if err == repo_errors.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		log.Println(err)
		return nil, ErrInternalServer
	}
	products := []dto.SaveProduct{}
	for _, p := range dbProducts {
		products = append(products, dto.SaveProduct{
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return products, nil
}

func (serv *productsServiceImpl) SaveProduct(ctx context.Context, product dto.SaveProduct) (string, error) {
	if err := serv.validate.StructCtx(ctx, product); err != nil {
		return "", ErrInvalidProductObject
	}
	tx, generatedId, err := serv.productsRepo.InsertProduct(ctx, models.Product{
		Name:        product.Name,
		Description: product.Description,
	})
	if err != nil {
		log.Println(err)
		return "", ErrInternalServer
	}
	return generatedId, tx.Commit()
}

func (serv *productsServiceImpl) UpdateProduct(ctx context.Context, product dto.UpdateProduct) error {
	if err := serv.validate.StructCtx(ctx, product); err != nil {
		return ErrInvalidProductObject
	}
	_, err := serv.GetProductById(ctx, product.Id)
	if err != nil {
		return err
	}
	tx, err := serv.productsRepo.UpdateProduct(ctx, models.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
	})
	if err != nil {
		log.Println(err)
		return ErrInternalServer
	}
	return tx.Commit()
}

func (serv *productsServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	if _, err := serv.GetProductById(ctx, id); err != nil {
		return err
	}
	tx, err := serv.productsRepo.DeleteProduct(ctx, id)
	if err != nil {
		log.Println(err)
		return ErrInternalServer
	}
	return tx.Commit()
}
