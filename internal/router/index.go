package router

import (
	"net/http"

	"example.com/ecommerce/internal/middlewares"
	"example.com/ecommerce/internal/modules/auth"
	"example.com/ecommerce/internal/modules/product"
	"github.com/gin-gonic/gin"
)

func Index() {
	initRoute := gin.Default()

	router := initRoute.Group("/api/v1")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Server Running",
		})
	})

	// Public Routes
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	//Protected Routes
	authorizedRoute := router.Group("", middlewares.AuthRequired())
	{
		authorizedRoute.POST("/create-product", product.CreateProduct)
		authorizedRoute.GET("/products", product.GetProducts)
		authorizedRoute.GET("/product/:id", product.GetProductByID)
		authorizedRoute.DELETE("/product/:id", product.DeleteProduct)
	}

	initRoute.Run(":8080")
}
