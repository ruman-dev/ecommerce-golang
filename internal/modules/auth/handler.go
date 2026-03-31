package auth

import (
	"net/http"

	"example.com/ecommerce/internal/modules/product"
	"example.com/ecommerce/models"
	"example.com/ecommerce/response"
	"example.com/ecommerce/validation"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, []string{"invalid request format"})
		return
	}

	// Validate
	if errors := validation.ValidateRegister(req); len(errors) > 0 {
		response.SendError(c, http.StatusBadRequest, errors)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, []string{"failed to hash password"})
		return
	}

	id, err := CreateUser(&req, string(hashedPassword))
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	user := models.User{
		ID:    id,
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}

	response.SendSuccess(c, http.StatusCreated, "user registered successfully", user)
}

func Login(c *gin.Context) {
	var req models.LoginRequest

	// Validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, []string{"invalid request format"})
		return
	}

	// Validate email and password format
	if errors := product.ValidateLogin(req); len(errors) > 0 {
		response.SendError(c, http.StatusBadRequest, errors)
		return
	}

	// Authenticate user
	user, token, err := LoginUser(req.Email, req.Password)
	if err != nil {
		response.SendError(c, http.StatusUnauthorized, []string{err.Error()})
		return
	}

	// Remove password from response
	user.Password = ""

	// Send success response
	response.SendSuccess(c, http.StatusOK, "login successful", models.LoginResponse{
		Token: token,
		User:  *user,
	})
}
