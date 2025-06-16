package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    string             `json:"uid" bson:"uid"`
	State     string             `json:"state" bson:"state"`
	Money     int                `json:"money" bson:"money"`
	CurrentAt time.Time          `json:"time" bson:"time"`
}
