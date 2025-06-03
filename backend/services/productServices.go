package services

import (
	"context"
	"errors"
	"time"

	"backend/config"
	"backend/models"

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
func GetSellingByID(uid string) ([]models.Product, error) {
	collection := config.GetCollection("Product")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	cursor, err := collection.Find(ctx, bson.M{"uid": uid})
	if err != nil {
		return []models.Product{}, err
	}
	var items []models.Product
	for cursor.Next(ctx) {
		var item models.Product
		err = cursor.Decode(&item)
		if err != nil {
			return []models.Product{}, err
		}
		items = append(items, item)
	}
	return items, nil
}

func CreateItem(item models.Product) (models.Product, error) {
	collection := config.GetCollection("Product")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	_, err := collection.InsertOne(ctx, item)
	if err != nil {
		return models.Product{}, err
	}
	return item, nil
}
func EditItem(item models.Product) (models.Product, error) {
	collection := config.GetCollection("Product")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	objID, err := primitive.ObjectIDFromHex(item.ID)
	if err != nil {
		return models.Product{}, errors.New("invalid item ID format")
	}

	updateField := bson.M{}
	if item.Name != "" {
		updateField["name"] = item.Name
	}
	if item.Description != "" {
		updateField["description"] = item.Description
	}
	if item.Price >= 0 {
		updateField["price"] = item.Price
	}
	if item.Quantity >= 0 {
		updateField["quantity"] = item.Quantity
	}
	if item.Category != nil {
		updateField["category"] = item.Category
	}
	if item.URL != "" {
		updateField["url"] = item.URL
	}
	var product models.Product
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, bson.M{"$set": updateField}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
func DeleteItem(pid string) error {
	collection := config.GetCollection("Product")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	objID, err := primitive.ObjectIDFromHex(pid)
	if err != nil {
		return errors.New("item ID not valid")
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
