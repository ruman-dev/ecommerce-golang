package product

import (
	"net/http"
	"strconv"

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

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	pId, err := strconv.Atoi(id)

	if err != nil {
		response.SendError(c, http.StatusBadRequest, []string{"Invalid ID Format"})
		return
	}

	for _, product := range products {
		if product.ID == pId {
			response.SendSuccess(c, http.StatusOK, "Product Returned Successfully", product)
			return
		}
	}

	response.SendError(c, http.StatusNotFound, []string{"Product Not Found"})

}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	pId, err := strconv.Atoi(id)

	if err != nil {
		response.SendError(c, http.StatusBadRequest, []string{"Invalid ID Format"})
		return
	}

	for i, product := range products {
		if product.ID == pId {
			products = append(products[:i], products[i+1:]...)
			response.SendSuccess(c, http.StatusOK, "Product Deleted Successfully", product)
			return
		}
	}
	response.SendError(c, http.StatusNotFound, []string{"Product Not Found"})
}
