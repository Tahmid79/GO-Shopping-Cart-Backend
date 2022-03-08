package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goecom1/configs"
	"goecom1/models/product"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Current struct {
	ID string `json:"_id"`
	product.Product
}

func enterData() {
	filename := "data.json"
	plan, _ := ioutil.ReadFile(filename)
	var data []Current
	json.Unmarshal(plan, &data)

	fmt.Println(data)

	newData := []product.Product{}
	allData := []interface{}{}

	for _, item := range data {
		newProduct := product.Product{}

		newProduct.ID = primitive.NewObjectID()

		newProduct.Description = item.Description
		newProduct.Picture = item.Picture
		newProduct.Price = item.Price
		newProduct.Stock = item.Stock
		newProduct.Title = item.Title

		newData = append(newData, newProduct)

		newDoc := bson.M{
			"_id":         newProduct.ID,
			"title":       newProduct.Title,
			"description": newProduct.Description,
			"picture":     newProduct.Picture,
			"price":       newProduct.Price,
			"stock":       newProduct.Stock,
		}

		allData = append(allData, newDoc)

		fmt.Println(newProduct.ID)
	}

	file, _ := json.MarshalIndent(newData, "", " ")

	_ = ioutil.WriteFile("newData.json", file, 0644)

	products := configs.GetCollection("products")

	result, err := products.InsertMany(context.TODO(), allData)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)

}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies(nil)

	configs.ConnectDB()
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
