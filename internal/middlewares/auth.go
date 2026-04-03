package middlewares

import (
	"net/http"
	"os"
	"strings"

	"example.com/ecommerce/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.SendError(c, http.StatusUnauthorized, []string{"authorization header required"})
			c.Abort()
			return
		}

		// Bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.SendError(c, http.StatusUnauthorized, []string{"invalid token format"})
			c.Abort()
			return
		}

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			response.SendError(c, http.StatusUnauthorized, []string{"invalid or expired token"})
			c.Abort()
			return
		}

		// Set user info in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
		}

		c.Next()
	}
}
