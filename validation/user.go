package validation

import "example.com/ecommerce/models"

func ValidateRegister(req models.RegisterRequest) []string {
	var errors []string

	if req.Name == "" {
		errors = append(errors, "name is required")
	}

	if req.Age <= 0 {
		errors = append(errors, "age must be greater than 0")
	}

	if req.Email == "" {
		errors = append(errors, "email is required")
	}

	if req.Password == "" {
		errors = append(errors, "password is required")
	} else if len(req.Password) < 6 {
		errors = append(errors, "password must be at least 6 characters")
	}

	return errors
}
