package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `json:"ID" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Duration    float64            `json:"duration" bson:"duration"`
	Year        int                `json:"year" bson:"year" `
	Description string             `json:"description" bson:"description"`
	Photo       string             `json:"photo" bson:"photo"`
}
