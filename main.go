package main

import (
	"goecom1/configs"
	"goecom1/models/product"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.ConnectDB()
	AddRoutes(router)
	router.Run("127.0.0.1:3000")
}

func AddRoutes(router *gin.Engine) {
	product.AddProductRoutes(router)
}
