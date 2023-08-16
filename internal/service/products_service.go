package service

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/otaxhu/go-htmx-project/internal/models"
	"github.com/otaxhu/go-htmx-project/internal/models/dto"
	"github.com/otaxhu/go-htmx-project/internal/repository"
	repo_errors "github.com/otaxhu/go-htmx-project/internal/repository/errors"
)

type ProductsService interface {
	GetProducts(ctx context.Context, page uint) ([]dto.GetProduct, error)
	GetProductById(ctx context.Context, id string) (dto.GetProduct, error)
	SaveProduct(ctx context.Context, product dto.SaveProduct) error
	UpdateProduct(ctx context.Context, product dto.UpdateProduct) error
	DeleteProduct(ctx context.Context, id string) error
}

type productsServiceImpl struct {
	productsRepo repository.ProductsRepository
	imageRepo    repository.ImageRepository
	validate     *validator.Validate
}

func NewProductsService(productsRepo repository.ProductsRepository, imageRepo repository.ImageRepository) ProductsService {
	return &productsServiceImpl{
		productsRepo: productsRepo,
		imageRepo:    imageRepo,
		validate:     validator.New(),
	}
}

func (serv *productsServiceImpl) GetProducts(ctx context.Context, page uint) ([]dto.GetProduct, error) {
	if page == 0 {
		return nil, ErrInvalidPageParam
	}
	page--
	const limit = 5
	offset := limit * page
	products, err := serv.productsRepo.GetProducts(ctx, offset, limit)
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
			ImageUrl:    product.ImageUrl,
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
		ImageUrl:    product.ImageUrl,
	}, nil
}

func (serv *productsServiceImpl) SaveProduct(ctx context.Context, product dto.SaveProduct) error {
	if err := serv.validate.StructCtx(ctx, product); err != nil {
		return ErrInvalidProductObject
	}
	tx, err := serv.productsRepo.UpsertProduct(ctx, models.Product{
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    product.Image.Header.Get("image_url"),
	})
	if err != nil {
		log.Println(err)
		return ErrInternalServer
	}
	if err := serv.imageRepo.SaveImage(product.Image); err != nil {
		tx.Rollback()
		log.Println(err)
		return ErrInternalServer
	}
	return tx.Commit()
}

func (serv *productsServiceImpl) UpdateProduct(ctx context.Context, product dto.UpdateProduct) error {
	if err := serv.validate.StructCtx(ctx, product); err != nil {
		return ErrInvalidProductObject
	}
	dbProduct, err := serv.GetProductById(ctx, product.Id)
	if err != nil {
		return err
	}
	var imageUrl string
	if product.Image != nil {
		imageUrl = product.Image.Header.Get("image_url")
	} else {
		imageUrl = dbProduct.ImageUrl
	}
	tx, err := serv.productsRepo.UpsertProduct(ctx, models.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    imageUrl,
	})
	if err != nil {
		log.Println(err)
		return ErrInternalServer
	}
	if product.Image == nil {
		return tx.Commit()
	}
	if err := serv.imageRepo.SaveImage(product.Image); err != nil {
		tx.Rollback()
		log.Println(err)
		return ErrInternalServer
	}
	return tx.Commit()
}

func (serv *productsServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	dbProduct, err := serv.GetProductById(ctx, id)
	if err != nil {
		return err
	}
	tx, err := serv.productsRepo.DeleteProduct(ctx, id)
	if err != nil {
		log.Println(err)
		return ErrInternalServer
	}
	if err := serv.imageRepo.DeleteImage(dbProduct.ImageUrl); err != nil {
		tx.Rollback()
		return ErrInternalServer
	}
	return tx.Commit()
}
