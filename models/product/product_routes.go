package product

import "github.com/gin-gonic/gin"

func AddProductRoutes(router *gin.Engine) {
	router.GET("/products", GetAllProducts())
}
