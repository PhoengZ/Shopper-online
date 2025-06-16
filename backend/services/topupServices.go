package services

import (
	"backend/config"
	"backend/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeRequest(userID string, amount int) error {
	collection := config.GetCollection("TopUp")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	transaction := models.Transaction{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		State:     "Approved",
		Money:     amount,
		CurrentAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}

func GetHistory() ([]map[string]interface{}, error) {
	collection := config.GetCollection("TopUp")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	var Transaction []map[string]interface{}
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []map[string]interface{}{}, err
	}
	for cursor.Next(ctx) {
		var item map[string]interface{}
		err := cursor.Decode(&item)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		if oid, ok := item["_id"].(primitive.ObjectID); ok {
			item["_id"] = oid.Hex()
		}
		Transaction = append(Transaction, item)
	}
	return Transaction, nil
}
