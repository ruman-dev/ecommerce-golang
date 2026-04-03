package product

import (
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
		return nil, err
	}

	return products, nil
}
