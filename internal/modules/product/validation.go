package product

import (
	"regexp"
	"strings"

	"example.com/ecommerce/models"
)

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
func ValidateLogin(req models.LoginRequest) []string {
	var errors []string

	// Validate email
	if req.Email == "" {
		errors = append(errors, "email is required")
	} else if !isValidEmail(req.Email) {
		errors = append(errors, "invalid email format")
	}

	// Validate password
	if req.Password == "" {
		errors = append(errors, "password is required")
	} else if len(req.Password) < 6 {
		errors = append(errors, "password must be at least 6 characters")
	}

	return errors
}
func isValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	// Basic email validation
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}
