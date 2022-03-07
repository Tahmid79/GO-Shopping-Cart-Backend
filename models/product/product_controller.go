package product

import (
	"goecom1/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = configs.GetCollection("products")

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		products := []Product{}

		ctx := configs.GetContext()

		cursor, err := productCollection.Find(ctx, bson.M{})

		if err != nil {
			body := ProductResponse{Message: err.Error(), Data: nil}
			c.JSON(http.StatusInternalServerError, body)
			return
		}

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			product := Product{}
			cursor.Decode(&product)
			products = append(products, product)
		}

		response := ProductResponse{Message: "Product List", Data: products}

		c.JSON(http.StatusOK, response)

	}
}
