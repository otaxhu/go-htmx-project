package repository

import (
	"context"
	"fmt"

	"github.com/otaxhu/go-htmx-project/database"
	"github.com/otaxhu/go-htmx-project/internal/models"
	"github.com/otaxhu/go-htmx-project/internal/repository/implementations"
	"github.com/otaxhu/go-htmx-project/internal/wrappers"
	"github.com/otaxhu/go-htmx-project/settings"
)

//go:generate mockery --name ProductsRepository --structname MockProductsRepository
type ProductsRepository interface {
	GetProducts(ctx context.Context, offset, limit uint) ([]models.Product, error)
	GetProductById(ctx context.Context, id string) (models.Product, error)
	UpsertProduct(ctx context.Context, product models.Product) (wrappers.Tx, error)
	DeleteProduct(ctx context.Context, id string) (wrappers.Tx, error)
}

func NewProductsRepository(dbSettings settings.Database) (ProductsRepository, error) {
	switch dbSettings.Driver {
	case "mysql":
		conn, err := database.GetSqlConnection(dbSettings)
		if err != nil {
			return nil, err
		}
		return implementations.NewMysqlProductsRepository(conn), nil
	default:
		return nil, fmt.Errorf("the %s driver does not have a `ProductsRepository` implementation", dbSettings.Driver)
	}
}
