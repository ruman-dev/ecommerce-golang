package auth

import (
	"database/sql"
	"errors"
	"time"

	"example.com/ecommerce/internal/config"
	"example.com/ecommerce/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req *models.RegisterRequest, hashedPassword string) (string, error) {
	query := `
        INSERT INTO users (name, age, email, password)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	var id string
	err := config.DB.QueryRow(query,
		req.Name,
		req.Age,
		req.Email,
		hashedPassword,
	).Scan(&id)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return "", errors.New("user already exists")
			}
		}
		return "", err
	}

	return id, nil
}

// Add this function to your existing auth package
func LoginUser(email, password string) (*models.User, string, error) {
	// Get user by email
	user, err := GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "User not found", err
		}
		return nil, "", err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := GenerateJWT(user)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := `SELECT * FROM users WHERE email=$1`

	err := config.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Generate JWT token for authenticated user
func GenerateJWT(user *models.User) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"age":     user.Age,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":     time.Now().Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	jwtSecret := []byte("JWT_SECRET") // Make sure to load from config
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify and parse JWT token
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	jwtSecret := []byte("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	return token, err
}
