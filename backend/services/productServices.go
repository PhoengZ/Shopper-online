package services

import (
	"context"
	"errors"
	"time"

	"backend/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(itemName string, category []string, price bool) ([]map[string]interface{}, error) {
	collection := config.GetCollection("Product")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	var product []map[string]interface{}
	var categoryBsonA bson.A
	for _, c := range category {
		categoryBsonA = append(categoryBsonA, c)
	}
	filter := bson.D{
		{Key: "name", Value: bson.D{
			{Key: "$regex", Value: itemName},
			{Key: "$options", Value: "i"},
		},
		},
	}
	if len(category) > 0 {
		filter = append(filter, bson.E{
			Key: "category", Value: bson.D{
				{Key: "$all", Value: categoryBsonA},
			},
		})
	}
	sortPrice := options.Find()
	if !price {
		sortPrice.SetSort(bson.D{{Key: "price", Value: -1}})
	} else {
		sortPrice.SetSort(bson.D{{Key: "price", Value: 1}})
	}
	cursor, err := collection.Find(ctx, filter, sortPrice)
	if err != nil {
		return product, err
	}
	for cursor.Next(ctx) {
		var prod map[string]interface{}
		err = cursor.Decode(&prod)
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
func GetSellingByID() {

}

func CreateItem() {

}
func EditItem() {

}
