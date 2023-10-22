package service

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/internal/models"
	"github.com/otaxhu/go-htmx-project/internal/models/dto"
	repo_errors "github.com/otaxhu/go-htmx-project/internal/repository/errors"
	repo_mocks "github.com/otaxhu/go-htmx-project/internal/repository/mocks"
	wrappers_mocks "github.com/otaxhu/go-htmx-project/internal/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Working mocks
var (
	productsRepoMock *repo_mocks.ProductsRepository
	txMock           *wrappers_mocks.Tx
	imageRepoMock    *repo_mocks.ImageRepository
)

// Can fail mocks
var (
	failingProductsRepoMock *repo_mocks.ProductsRepository
	failingImageRepoMock    *repo_mocks.ImageRepository
)

// Working ProductsService for testing
var (
	productsService ProductsService
)

// Can fail ProductsService for testing
var (
	failingProductsRepoProductsService ProductsService
	failingImageRepoProductsService    ProductsService
)

var (
	validUuid    string
	notExistUuid string
)

func TestMain(m *testing.M) {
	productsRepoMock = &repo_mocks.ProductsRepository{}
	txMock = &wrappers_mocks.Tx{}
	imageRepoMock = &repo_mocks.ImageRepository{}
	cfgProductsService := config.ProductsService{GetProductsLimit: 50}
	productsService = NewProductsService(cfgProductsService, productsRepoMock)
	validUuid = uuid.NewString()
	notExistUuid = uuid.NewString()
	failingProductsRepoMock = &repo_mocks.ProductsRepository{}
	failingImageRepoMock = &repo_mocks.ImageRepository{}
	failingProductsRepoProductsService = NewProductsService(cfgProductsService, failingProductsRepoMock)
	failingImageRepoProductsService = NewProductsService(cfgProductsService, productsRepoMock)
	txMock.On("Commit").Return(nil)
	txMock.On("Rollback").Return(nil)
	productsRepoMock.On("DeleteProductById", mock.Anything, mock.Anything).Return(txMock, nil)
	productsRepoMock.On("GetProductById", mock.Anything, validUuid).Return(models.Product{}, nil)
	productsRepoMock.On("GetProductById", mock.Anything, notExistUuid).Return(models.Product{}, repo_errors.ErrNoRows)
	productsRepoMock.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).Return([]models.Product{{}}, nil)
	productsRepoMock.On("InsertProduct", mock.Anything, mock.Anything).Return(txMock, validUuid, nil)
	productsRepoMock.On("UpdateProduct", mock.Anything, mock.Anything).Return(txMock, nil)
	imageRepoMock.On("DeleteImage", mock.Anything).Return(nil)
	imageRepoMock.On("SaveImage", mock.Anything).Return(nil)
	failingProductsRepoMock.On("DeleteProductById", mock.Anything, mock.Anything).Return(nil, errors.New("test error"))
	failingProductsRepoMock.On("GetProductById", mock.Anything, mock.Anything).Return(models.Product{}, errors.New("test error"))
	failingProductsRepoMock.On("GetProducts", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("test error"))
	failingProductsRepoMock.On("InsertProduct", mock.Anything, mock.Anything).Return(nil, "", errors.New("test error"))
	failingProductsRepoMock.On("UpdateProduct", mock.Anything, mock.Anything).Return(nil, errors.New("test error"))
	failingImageRepoMock.On("DeleteImage", mock.Anything).Return(errors.New("test error"))
	failingImageRepoMock.On("SaveImage", mock.Anything).Return(errors.New("test error"))
	os.Exit(m.Run())
}

func TestGetProducts(t *testing.T) {
	testCasesWorking := []struct {
		name          string
		page          int
		expectedError error
	}{
		{
			name:          "GetProducts_Success",
			page:          1,
			expectedError: nil,
		},
		{
			name:          "GetProducts_InvalidPageParam",
			page:          0,
			expectedError: ErrInvalidInput,
		},
	}

	testCasesFailingProductsRepo := []struct {
		name          string
		expectedError error
	}{
		{
			name:          "GetProducts_FailingProductsRepo",
			expectedError: ErrInternalServer,
		},
	}

	ctx := context.Background()

	for i := range testCasesWorking {
		tc := testCasesWorking[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			productsRepoMock.Test(t)
			txMock.Test(t)
			imageRepoMock.Test(t)

			_, err := productsService.GetProducts(ctx, tc.page)
			assert.Equal(t, tc.expectedError, err)
		})
	}
	for i := range testCasesFailingProductsRepo {
		tc := testCasesFailingProductsRepo[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			imageRepoMock.Test(t)
			failingProductsRepoMock.Test(t)

			_, err := failingProductsRepoProductsService.GetProducts(ctx, 1)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetProductById(t *testing.T) {
	testCasesWorking := []struct {
		name          string
		id            string
		expectedError error
	}{
		{
			name:          "GetProductById_Success",
			id:            validUuid,
			expectedError: nil,
		},
		{
			name:          "GetProductById_NotFound",
			id:            notExistUuid,
			expectedError: ErrNotFound,
		},
	}
	testCasesFailingProductsRepo := []struct {
		name          string
		id            string
		expectedError error
	}{
		{
			name:          "GetProductById_FailingProductsRepo",
			id:            validUuid,
			expectedError: ErrInternalServer,
		},
	}

	ctx := context.Background()

	for i := range testCasesWorking {
		tc := testCasesWorking[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			productsRepoMock.Test(t)
			txMock.Test(t)
			imageRepoMock.Test(t)

			_, err := productsService.GetProductById(ctx, tc.id)
			assert.Equal(t, tc.expectedError, err)
		})
	}
	for i := range testCasesFailingProductsRepo {
		tc := testCasesFailingProductsRepo[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			imageRepoMock.Test(t)
			failingProductsRepoMock.Test(t)

			_, err := failingProductsRepoProductsService.GetProductById(ctx, tc.id)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestSaveProduct(t *testing.T) {
	testCasesWorking := []struct {
		name          string
		productData   dto.SaveProduct
		expectedError error
	}{
		{
			name: "SaveProduct_Success",
			productData: dto.SaveProduct{
				Name:        "test",
				Description: "test",
			},
			expectedError: nil,
		},
		{
			name: "SaveProduct_InvalidProductObject",
			productData: dto.SaveProduct{
				Name:        "",
				Description: "",
			},
			expectedError: ErrInvalidInput,
		},
	}

	testCasesFailingProductsRepo := []struct {
		name          string
		expectedError error
	}{
		{
			name:          "SaveProduct_FailingProductsRepo",
			expectedError: ErrInternalServer,
		},
	}

	ctx := context.Background()

	for i := range testCasesWorking {
		tc := testCasesWorking[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			productsRepoMock.Test(t)
			txMock.Test(t)
			imageRepoMock.Test(t)

			_, err := productsService.SaveProduct(ctx, tc.productData)
			assert.Equal(t, tc.expectedError, err)
		})
	}
	for i := range testCasesFailingProductsRepo {
		tc := testCasesFailingProductsRepo[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			imageRepoMock.Test(t)
			failingProductsRepoMock.Test(t)

			_, err := failingProductsRepoProductsService.SaveProduct(ctx, dto.SaveProduct{
				Name:        "test",
				Description: "test",
			})
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	testCasesWorking := []struct {
		name          string
		productData   dto.UpdateProduct
		expectedError error
	}{
		{
			name: "UpdateProduct_Success",
			productData: dto.UpdateProduct{
				Id:          validUuid,
				Name:        "test",
				Description: "test",
			},
			expectedError: nil,
		},
		{
			name: "UpdateProduct_InvalidProductObject",
			productData: dto.UpdateProduct{
				Id:          "this is not an UUID",
				Name:        "test",
				Description: "test",
			},
			expectedError: ErrInvalidInput,
		},
		{
			name: "UpdateProduct_NotFound",
			productData: dto.UpdateProduct{
				Id:          notExistUuid,
				Name:        "test",
				Description: "test",
			},
			expectedError: ErrNotFound,
		},
		{
			name: "UpdateProduct_Success_NoImage_UsingTheDbImage",
			productData: dto.UpdateProduct{
				Id:          validUuid,
				Name:        "test",
				Description: "test",
			},
			expectedError: nil,
		},
	}

	testCasesFailingProductsRepo := []struct {
		name          string
		expectedError error
	}{
		{
			name:          "UpdateProduct_FailingProductsRepo",
			expectedError: ErrInternalServer,
		},
	}

	ctx := context.Background()

	for i := range testCasesWorking {
		tc := testCasesWorking[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			productsRepoMock.Test(t)
			txMock.Test(t)
			imageRepoMock.Test(t)

			_, err := productsService.UpdateProduct(ctx, tc.productData)
			assert.Equal(t, tc.expectedError, err)
		})
	}
	for i := range testCasesFailingProductsRepo {
		tc := testCasesFailingProductsRepo[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			failingProductsRepoMock.Test(t)
			imageRepoMock.Test(t)

			_, err := failingProductsRepoProductsService.UpdateProduct(ctx, dto.UpdateProduct{
				Id:          validUuid,
				Name:        "test",
				Description: "test",
			})
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	testCasesWorking := []struct {
		name          string
		id            string
		expectedError error
	}{
		{
			name:          "DeleteProduct_Success",
			id:            validUuid,
			expectedError: nil,
		},
		{
			name:          "DeleteProduct_NotFound",
			id:            notExistUuid,
			expectedError: ErrNotFound,
		},
	}

	testCasesFailingProductsRepo := []struct {
		name          string
		expectedError error
	}{
		{
			name:          "DeleteProduct_FailingProductsRepo",
			expectedError: ErrInternalServer,
		},
	}

	ctx := context.Background()

	for i := range testCasesWorking {
		tc := testCasesWorking[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			productsRepoMock.Test(t)
			txMock.Test(t)
			imageRepoMock.Test(t)

			err := productsService.DeleteProductById(ctx, tc.id)
			assert.Equal(t, tc.expectedError, err)
		})
	}
	for i := range testCasesFailingProductsRepo {
		tc := testCasesFailingProductsRepo[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			failingProductsRepoMock.Test(t)
			imageRepoMock.Test(t)

			err := failingProductsRepoProductsService.DeleteProductById(ctx, validUuid)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
