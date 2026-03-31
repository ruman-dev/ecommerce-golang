package response

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status bool     `json:"status"`
	Errors []string `json:"errors"`
}

func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
	c.IndentedJSON(statusCode, response)
}

func SendError(c *gin.Context, statusCode int, errors []string) {
	response := ErrorResponse{
		Status: false,
		Errors: errors,
	}
	c.IndentedJSON(statusCode, response)
}
