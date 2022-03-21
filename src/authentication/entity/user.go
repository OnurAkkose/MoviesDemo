package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"ID" bson:"_id"`
	FullName string             `json:"username" bson:"username"`
}
