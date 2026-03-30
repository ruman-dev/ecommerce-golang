package product

import "example.com/ecommerce/models"

func ValidateCreateProduct(p models.Product) []string {
	var errors []string

	if p.Title == "" {
		errors = append(errors, "title is required")
	}

	if p.Description == "" {
		errors = append(errors, "description is required")
	}

	if p.Amount <= 0 {
		errors = append(errors, "amount must be greater than 0")
	}

	if p.Quantity < 0 {
		errors = append(errors, "quantity cannot be negative")
	}

	return errors
}
