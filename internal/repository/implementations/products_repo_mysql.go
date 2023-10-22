package implementations

import (
	"context"
	"database/sql"
	"log"
	"strings"

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
	qryGetProducts    = "SELECT id, name, description FROM products LIMIT ? OFFSET ?"
	qryGetProductById = "SELECT id, name, description FROM products WHERE id = ?"
	qrySearchProducts = `SELECT
							id,
							name,
							description
						FROM products
						WHERE
							-- Exact coincidence
							name = ? OR
							description = ? OR

							-- Starts with ?
							name LIKE ? OR
							description LIKE ? OR

							-- Ends with ?
							name LIKE ? OR
							description LIKE ? OR

							-- Inner match ?
							name LIKE ? OR
							description LIKE ?`
	qryInsertProduct = "INSERT INTO products (id, name, description) VALUES (?, ?, ?)"
	qryDeleteProduct = "DELETE FROM products WHERE id = ?"
	qryUpdateProduct = "UPDATE products SET name = ?, description = ? WHERE id = ?"
)

func (repo *mysqlProductsRepo) Close() error {
	return repo.db.Close()
}

func (repo *mysqlProductsRepo) GetProducts(ctx context.Context, offset, limit int) ([]models.Product, error) {
	rows, err := repo.db.QueryContext(ctx, qryGetProducts, limit, offset)
	if err != nil {
		return nil, err
	}
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.Id, &product.Name, &product.Description); err != nil {
			log.Println(err)
			continue
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
	if err := repo.db.QueryRowContext(ctx, qryGetProductById, id).Scan(&product.Id, &product.Name, &product.Description); err == sql.ErrNoRows {
		return product, repo_errors.ErrNoRows
	} else if err != nil {
		return product, err
	}
	return product, nil
}

func sanatizeLikeTerm(term string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(term, "\\", "\\\\"), "%", "\\%"), "_", "\\_")
}

func (repo *mysqlProductsRepo) SearchProducts(ctx context.Context, term string) ([]models.Product, error) {
	sanatizedLikeTerm := sanatizeLikeTerm(term)
	sanatizedStartsWithTerm := sanatizedLikeTerm + "%"
	sanatizedEndsWithTerm := "%" + sanatizedLikeTerm
	sanatizedInnerMatchTerm := "%" + sanatizedLikeTerm + "%"
	rows, err := repo.db.QueryContext(ctx, qrySearchProducts,
		term,
		term,
		sanatizedStartsWithTerm,
		sanatizedStartsWithTerm,
		sanatizedEndsWithTerm,
		sanatizedEndsWithTerm,
		sanatizedInnerMatchTerm,
		sanatizedInnerMatchTerm,
	)
	if err != nil {
		return nil, err
	}
	products := []models.Product{}
	for rows.Next() {
		p := models.Product{}
		if err := rows.Scan(&p.Id, &p.Name, &p.Description); err != nil {
			log.Println(err)
			continue
		}
		products = append(products, p)
	}
	if len(products) == 0 {
		return nil, repo_errors.ErrNoRows
	}
	return products, nil
}

func (repo *mysqlProductsRepo) InsertProduct(ctx context.Context, product models.Product) (wrappers.Tx, string, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, "", err
	}
	generatedId := uuid.NewString()
	if _, err := tx.ExecContext(ctx, qryInsertProduct, generatedId, product.Name, product.Description); err != nil {
		tx.Rollback()
		return nil, "", err
	}
	return tx, generatedId, nil
}

func (repo *mysqlProductsRepo) UpdateProduct(ctx context.Context, product models.Product) (wrappers.Tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if _, err := tx.ExecContext(ctx, qryUpdateProduct, product.Name, product.Description, product.Id); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tx, nil
}

func (repo *mysqlProductsRepo) DeleteProductById(ctx context.Context, id string) (wrappers.Tx, error) {
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
