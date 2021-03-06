package main

import (
	"goecom1/models/product"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	mode := os.Getenv("GIN_MODE")
	gin.SetMode(mode)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(CORSMiddleware())
	AddRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := router.Run("localhost:" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func AddRoutes(router *gin.Engine) {
	product.AddProductRoutes(router)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
