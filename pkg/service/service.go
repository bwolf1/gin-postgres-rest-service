package service

import (
	"context"
	"fmt"

	"github.com/bwolf1/gin-postgres-rest-service/pkg/model"
)

// Create and interface for the product
type ProductInterface interface {
	ListProducts(ctx context.Context, productName string, page, pageSize int, sort string) ([]model.Product, error)
	GetProduct(ctx context.Context, id string) (model.Product, error)
	GetProductVersions(ctx context.Context, id string) ([]model.ProductVersion, error)
}

// ListProducts returns a list of products
func (p *Product) ListProducts(
	ctx context.Context, productName string, page, pageSize int, sort string,
) ([]model.Product, error) {
	// TODO: Add input validation so that only name, start, page_size, and sort are allowed

	// Ensure that start is not negative and reflects the correct page
	if page <= 0 {
		page = 0
	} else {
		page--
	}
	query := fmt.Sprintf(
		"SELECT id, name, description, version_count FROM product WHERE name LIKE $1 ORDER BY %s LIMIT $2 OFFSET $3",
		sort,
	)
	rows, err := p.db.Query(query, "%"+productName+"%", pageSize, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.VersionCount)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProduct returns a product by ID
func (p *Product) GetProduct(ctx context.Context, id string) (model.Product, error) {
	row := p.db.QueryRow("SELECT id, name, description, version_count FROM product WHERE id = $1", id)
	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.VersionCount)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

// GetProductVersions returns a list of product versions by product ID
func (p *Product) GetProductVersions(ctx context.Context, id string) ([]model.ProductVersion, error) {
	rows, err := p.db.Query(
		"SELECT id, product_id, version, description FROM product_version WHERE product_id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []model.ProductVersion
	for rows.Next() {
		var version model.ProductVersion
		err := rows.Scan(&version.ID, &version.ProductID, &version.Version, &version.Description)
		if err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return versions, nil
}
