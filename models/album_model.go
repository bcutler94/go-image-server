package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AlbumRequest struct {
	ID     primitive.ObjectID `json:"_id"`
	Title  string             `json:"title" validate:"required"`
	Artist string             `json:"artist" validate:"required"`
	Price  float64            `json:"price" validate:"required"`
}

type AlbumModel struct {
	Title  string  `bson:"title" validate:"required"`
	Artist string  `bson:"artist" validate:"required"`
	Price  float64 `bson:"price" validate:"required"`
}
