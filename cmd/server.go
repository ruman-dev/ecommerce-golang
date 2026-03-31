package cmd

import (
	"example.com/ecommerce/internal/config"
	"example.com/ecommerce/internal/router"
)

func Start() {
	config.ConnectDB()

	router.Index()
}
