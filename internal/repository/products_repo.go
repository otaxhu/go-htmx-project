package repository

import (
	"context"
	"fmt"

	"github.com/otaxhu/go-htmx-project/config"
	"github.com/otaxhu/go-htmx-project/database"
	"github.com/otaxhu/go-htmx-project/internal/models"
	"github.com/otaxhu/go-htmx-project/internal/repository/implementations"
	"github.com/otaxhu/go-htmx-project/internal/wrappers"
)

//go:generate mockery --name ProductsRepository
type ProductsRepository interface {
	GetProducts(ctx context.Context, offset, limit int) ([]models.Product, error)
	GetProductById(ctx context.Context, id string) (models.Product, error)
	SearchProducts(ctx context.Context, term string) ([]models.Product, error)
	InsertProduct(ctx context.Context, product models.Product) (wrappers.Tx, string, error)
	UpdateProduct(ctx context.Context, product models.Product) (wrappers.Tx, error)
	DeleteProduct(ctx context.Context, id string) (wrappers.Tx, error)
}

func NewProductsRepository(dbCfg config.Database) (ProductsRepository, error) {
	switch dbCfg.Driver {
	case "mysql":
		conn, err := database.GetSqlConnection(dbCfg)
		if err != nil {
			return nil, err
		}
		return implementations.NewMysqlProductsRepository(conn), nil
	default:
		return nil, fmt.Errorf("the %s driver does not have a `ProductsRepository` implementation", dbCfg.Driver)
	}
}
