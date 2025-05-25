package services

import (
	"context"
	"errors"
	"time"

	"backend/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts() ([]map[string]interface{}, error) {
	collection := config.GetCollection("Product")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	var product []map[string]interface{}
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return product, err
	}
	for cursor.Next(ctx) {
		var prod map[string]interface{}
		err := cursor.Decode(&prod)
		if err != nil {
			return []map[string]interface{}{}, err
		}
		product = append(product, prod)
	}
	return product, nil
}

func GetProductByID(productID string) (map[string]interface{}, error) {
	collection := config.GetCollection("Product")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}
	var product map[string]interface{}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return product, nil
}
