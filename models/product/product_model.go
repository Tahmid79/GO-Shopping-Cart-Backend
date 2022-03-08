package product

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id""`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Picture     string             `json:"picture"`
	Price       int                `json:"price"`
	Stock       int                `json:"stock"`
}
