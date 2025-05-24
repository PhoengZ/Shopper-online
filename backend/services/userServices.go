package services

import (
	"context"
	"errors"
	"time"

	"backend/config"
	"backend/models"
	"backend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(username, password string) (string, error) {
	collection := config.GetCollection("User")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", errors.New("username not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func RegisterUser(user models.User) error {
	collection := config.GetCollection("User")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	var existUser models.User

	err := collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existUser)
	if err == nil {
		return errors.New("username already used")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	_, err = collection.InsertOne(ctx, user)
	return err
}

func GetCartListByID(id string) ([]map[string]interface{}, error) {
	collection := config.GetCollection("User")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	return user.CartList, nil
}

func addItemOnCart(id string, item map[string]interface{}) error {
	collection := config.GetCollection("User")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return err
	}
	productID, ok := item["id"].(string)
	if !ok {
		return errors.New("invalid to decode productID")
	}
	found := false
	var quantity int
	l := len(user.CartList)
	for i := 0; i < l; i++ {
		if user.CartList[i]["id"] == productID {
			found = true
			quantity = user.CartList[i]["quantity"].(int) + 1
			break
		}
	}
	if found {
		filter := bson.M{
			"_id":         objID,
			"cartlist.id": productID,
		}
		update := bson.M{
			"$inc": bson.M{
				"cartlist.$.quantity": quantity,
			},
		}
		_, err = collection.UpdateOne(ctx, filter, update)
	} else {
		item["quantity"] = 1
		update := bson.M{
			"$push": bson.M{
				"cartlist": item,
			},
		}
		_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	}
	if err != nil {
		return err
	}
	return nil
}
