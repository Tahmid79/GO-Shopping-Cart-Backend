package main

import (
	"goecom1/configs"
	"goecom1/models/product"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	configs.ConnectDB()
	AddRoutes(router)
	// router.Logger.Fatal(router.Start("localhost:3000"))

	if err := router.Start(":3000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

func AddRoutes(router *echo.Echo) {
	product.AddProductRoutes(router)
}
