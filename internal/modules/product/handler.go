package product

import (
	"net/http"
	"time"

	"example.com/ecommerce/models"
	"example.com/ecommerce/response"
	"github.com/gin-gonic/gin"
)

var products []models.Product

// func CreateProduct(c *gin.Context) {
// 	var p models.Product
// 	if err := c.BindJSON(&p); err != nil {
// 		response.SendError(c, http.StatusBadRequest, []string{"Invalid JSON payload"})
// 		return
// 	}

// 	if errors := ValidateCreateProduct(p); len(errors) > 0 {
// 		response.SendError(c, http.StatusBadRequest, errors)
// 		return
// 	}
// 	p.ID = len(products) + 1

//		products = append(products, p)
//		response.SendSuccess(c, http.StatusCreated, "Product Created Successfully", p)
//	}
func CreateProduct(c *gin.Context) {
	var p models.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		response.SendError(c, http.StatusBadRequest, []string{"Invalid JSON payload"})
		return
	}

	if errs := ValidateCreateProduct(p); len(errs) > 0 {
		response.SendError(c, http.StatusBadRequest, errs)
		return
	}

	// Save to DB
	id, err := CreateProductDB(&p)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	p.ID = id
	p.Created_At = time.Now()

	response.SendSuccess(c, http.StatusCreated, "Product Created Successfully", p)
}

func GetProducts(c *gin.Context) {

	products, err := GetProductsDB()

	if err != nil {
		response.SendError(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	if products == nil {
		response.SendError(c, http.StatusNotFound, []string{"No products found"})
		return
	}

	response.SendSuccess(c, http.StatusOK, "Products Returned Successfully", products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	product, err := GetProductByIdDB(id)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	// ✅ Handle not found
	if product == nil {
		response.SendError(c, http.StatusNotFound, []string{"Product not found"})
		return
	}

	response.SendSuccess(c, http.StatusOK, "Product fetched successfully", product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	_, err := DeleteProductDB(id)
	if err != nil {
		response.SendError(c, http.StatusNotFound, []string{err.Error()})
		return
	}
	response.SendSuccess(c, http.StatusOK, "Product deleted successfully", []string{})
}
