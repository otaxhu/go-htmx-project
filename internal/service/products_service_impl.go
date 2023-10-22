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

type productsServiceImpl struct {
	cfg          config.ProductsService
	productsRepo repository.ProductsRepository
	validate     *validator.Validate
}

func (serv *productsServiceImpl) GetProducts(ctx context.Context, page int) ([]dto.GetProduct, error) {
	if page <= 0 {
		return nil, ErrInvalidInput
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
	if err := serv.validate.VarCtx(ctx, id, "required,uuid"); err != nil {
		return dto.GetProduct{}, ErrInvalidInput
	}
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

func (serv *productsServiceImpl) SearchProducts(ctx context.Context, term string) ([]dto.GetProduct, error) {
	if err := serv.validate.VarCtx(ctx, term, "required"); err != nil {
		return nil, ErrInvalidInput
	}
	dbProducts, err := serv.productsRepo.SearchProducts(ctx, term)
	if err == repo_errors.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		log.Println(err)
		return nil, ErrInternalServer
	}
	products := []dto.GetProduct{}
	for _, p := range dbProducts {
		products = append(products, dto.GetProduct{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return products, nil
}

func (serv *productsServiceImpl) SaveProduct(ctx context.Context, product dto.SaveProduct) (dto.GetProduct, error) {
	if err := serv.validate.StructCtx(ctx, product); err != nil {
		return dto.GetProduct{}, ErrInvalidInput
	}
	tx, generatedId, err := serv.productsRepo.InsertProduct(ctx, models.Product{
		Name:        product.Name,
		Description: product.Description,
	})
	if err != nil {
		log.Println(err)
		return dto.GetProduct{}, ErrInternalServer
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return dto.GetProduct{}, ErrInternalServer
	}
	return serv.GetProductById(ctx, generatedId)
}

func (serv *productsServiceImpl) UpdateProduct(ctx context.Context, product dto.UpdateProduct) (dto.GetProduct, error) {
	if err := serv.validate.StructCtx(ctx, product); err != nil {
		return dto.GetProduct{}, ErrInvalidInput
	}
	dbProduct, err := serv.GetProductById(ctx, product.Id)
	if err != nil {
		return dto.GetProduct{}, err
	}
	if dbProduct == dto.GetProduct(product) {
		log.Println("SON IGUALES LOS PRODUCTOS ACTUALIZADOS!!!!!!")
		return dbProduct, nil
	}
	tx, err := serv.productsRepo.UpdateProduct(ctx, models.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
	})
	if err != nil {
		log.Println(err)
		return dto.GetProduct{}, ErrInternalServer
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return dto.GetProduct{}, ErrInternalServer
	}
	return serv.GetProductById(ctx, product.Id)
}

func (serv *productsServiceImpl) DeleteProductById(ctx context.Context, id string) error {
	if _, err := serv.GetProductById(ctx, id); err != nil {
		return err
	}
	tx, err := serv.productsRepo.DeleteProductById(ctx, id)
	if err != nil {
		log.Println(err)
		return ErrInternalServer
	}
	return tx.Commit()
}
