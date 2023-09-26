package implementations

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/otaxhu/go-htmx-project/internal/models"
	repo_errors "github.com/otaxhu/go-htmx-project/internal/repository/errors"
	"github.com/otaxhu/go-htmx-project/internal/wrappers"
)

type mysqlProductsRepo struct {
	db *sql.DB
}

func NewMysqlProductsRepository(db *sql.DB) *mysqlProductsRepo {
	return &mysqlProductsRepo{
		db: db,
	}
}

const (
	qryGetProducts    = "SELECT id, name, description, image_url FROM products LIMIT ? OFFSET ?"
	qryGetProductById = "SELECT id, name, description, image_url FROM products WHERE id = ?"
	qryInsertProduct  = "INSERT INTO products (id, name, description, image_url) VALUES (?, ?, ?, ?)"
	qryDeleteProduct  = "DELETE FROM products WHERE id = ?"
	qryUpdateProduct  = "UPDATE products SET name = ?, description = ?, image_url = ? WHERE id = ?"
)

func (repo *mysqlProductsRepo) GetProducts(ctx context.Context, offset, limit uint) ([]models.Product, error) {
	rows, err := repo.db.QueryContext(ctx, qryGetProducts, limit, offset)
	if err != nil {
		return nil, err
	}
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if len(products) == 0 {
		return nil, repo_errors.ErrNoRows
	}
	return products, nil
}

func (repo *mysqlProductsRepo) GetProductById(ctx context.Context, id string) (models.Product, error) {
	product := models.Product{}
	if err := repo.db.QueryRowContext(ctx, qryGetProductById, id).Scan(&product.Id, &product.Name, &product.Description, &product.ImageUrl); err == sql.ErrNoRows {
		return product, repo_errors.ErrNoRows
	} else if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *mysqlProductsRepo) UpsertProduct(ctx context.Context, product models.Product) (wrappers.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if product.Id != "" {
		if _, err := repo.GetProductById(ctx, product.Id); err != nil && err != repo_errors.ErrNoRows {
			tx.Rollback()
			return nil, err
		} else if err == repo_errors.ErrNoRows {
			// Insert if there is no record in the database
			if _, err := tx.ExecContext(ctx, qryInsertProduct, uuid.NewString(), product.Name, product.Description, product.ImageUrl); err != nil {
				tx.Rollback()
				return nil, err
			}
			return tx, nil
		}
		// Update if there is record in the database
		if _, err := tx.ExecContext(ctx, qryUpdateProduct, product.Name, product.Description, product.ImageUrl, product.Id); err != nil {
			tx.Rollback()
			return nil, err
		}
		return tx, nil
	}

	// Insert if there is no Id field in the product struct
	if _, err := tx.ExecContext(ctx, qryInsertProduct, uuid.NewString(), product.Name, product.Description, product.ImageUrl); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tx, nil
}

func (repo *mysqlProductsRepo) DeleteProduct(ctx context.Context, id string) (wrappers.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if _, err := tx.ExecContext(ctx, qryDeleteProduct, id); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tx, nil
}
