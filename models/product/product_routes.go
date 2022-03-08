package product

import (
	"github.com/labstack/echo/v4"
)

func AddProductRoutes(router *echo.Echo) {
	router.GET("/products", GetAllProducts())
}
