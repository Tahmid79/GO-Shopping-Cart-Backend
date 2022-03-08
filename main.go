package main

import (
	"goecom1/configs"
	"goecom1/models/product"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	configs.ConnectDB()
	AddRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func AddRoutes(router *gin.Engine) {
	product.AddProductRoutes(router)
}
