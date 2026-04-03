package product

import (
	"database/sql"

	"example.com/ecommerce/internal/config"
	"example.com/ecommerce/models"
)

func CreateProductDB(p *models.Product) (int, error) {
	query := `
		INSERT INTO products (name, price, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int

	err := config.DB.QueryRow(
		query,
		p.Name,
		p.Price,
		p.Description,
	).Scan(&id)

	return id, err
}

func GetProductsDB() ([]models.Product, error) {
	var products []models.Product

	query := `SELECT * FROM products`

	err := config.DB.Select(&products, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return products, nil
}

func DeleteProductDB(id string) (*models.Product, error) {
	var product models.Product

	query := `DELETE FROM products WHERE id=$1`

	err := config.DB.Get(&product, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}
func GetProductByIdDB(id string) (*models.Product, error) {
	var product models.Product

	query := `SELECT id, name, price, quantity, description FROM products WHERE id=$1`

	err := config.DB.Get(&product, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}
