package main

import (
	"goecom1/configs"
	"goecom1/models/product"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	configs.ConnectDB()
	AddRoutes(router)
	router.Logger.Fatal(router.Start("localhost:3000"))
}

func AddRoutes(router *echo.Echo) {
	product.AddProductRoutes(router)
}
