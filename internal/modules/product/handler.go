package product

import (
	"net/http"

	"example.com/ecommerce/models"
	"example.com/ecommerce/response"
	"github.com/gin-gonic/gin"
)

var products []models.Product

func CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.BindJSON(&p); err != nil {
		response.SendError(c, http.StatusBadRequest, []string{"Invalid JSON payload"})
		return
	}

	if errors := ValidateCreateProduct(p); len(errors) > 0 {
		response.SendError(c, http.StatusBadRequest, errors)
		return
	}
	p.ID = len(products) + 1

	products = append(products, p)
	response.SendSuccess(c, http.StatusCreated, "Product Created Successfully", p)
}

func GetProducts(c *gin.Context) {
	response.SendSuccess(c, http.StatusOK, "Products Returned Successfully", products)
}
