package auth

import (
	"net/http"
	"time"

	"example.com/ecommerce/models"
	"example.com/ecommerce/response"
	"example.com/ecommerce/validation"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var users []models.User
var jwtSecret = []byte("your-secret-key") // In production, use environment variable

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

	// Create user
	user := models.User{
		ID:       len(users) + 1,
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	users = append(users, user)

	response.SendSuccess(c, http.StatusCreated, "user registered successfully", user)
}

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendError(c, http.StatusBadRequest, []string{"invalid request format"})
		return
	}

	// Find user
	var user models.User
	found := false
	for _, u := range users {
		if u.Email == req.Email {
			user = u
			found = true
			break
		}
	}

	if !found {
		response.SendError(c, http.StatusUnauthorized, []string{"invalid credentials"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response.SendError(c, http.StatusUnauthorized, []string{"invalid credentials"})
		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, []string{"failed to generate token"})
		return
	}

	response.SendSuccess(c, http.StatusOK, "Login successful", gin.H{"token": tokenString})
}
