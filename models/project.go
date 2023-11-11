package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `bson:"id,omitempty"`
	Name        string             `json:"name"`
	Intro       string             `json:"intro"`
	Image       string             `json:"image"`
	Description []string           `json:"description"`
	Items       []Item             `json:"items"`
	Images      []Image            `json:"images,omitempty"`
}

type Item struct {
	Title string   `json:"title"`
	Value []string `json:"value"`
}

type Image struct {
	Title string `json:"title"`
	Image string `json:"image"`
}
