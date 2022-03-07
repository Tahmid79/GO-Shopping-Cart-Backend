package main

import (
	"goecom1/configs"
	"goecom1/models/product"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	configs.ConnectDB()
	AddRoutes(router)
	router.Run("localhost:3000")
}

func AddRoutes(router *gin.Engine) {
	product.AddProductRoutes(router)
}
