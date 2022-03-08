package product

import (
	"goecom1/configs"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection("products")

func GetAllProducts() echo.HandlerFunc {
	return func(c echo.Context) error {

		products := []Product{}

		ctx := configs.GetContext()

		cursor, err := productCollection.Find(ctx, bson.M{})

		if err != nil {
			body := ProductResponse{Message: err.Error(), Data: nil}
			return c.JSON(http.StatusInternalServerError, body)
		}

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			product := Product{}
			cursor.Decode(&product)
			products = append(products, product)
		}

		response := ProductResponse{Message: "Product List", Data: products}

		return c.JSON(http.StatusOK, response)
	}
}
