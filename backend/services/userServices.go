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

func GetCartListByID(id string) ([]models.Item, error) {
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
		return []models.Item{}, err
	}
	return user.CartList, nil
}

func RemoveItemOnCart(userID, itemID string) error {
	collection := config.GetCollection("User")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return err
	}
	l := len(user.CartList)
	found := false
	for i := 0; i < l; i++ {
		if user.CartList[i].ID == itemID {
			found = true
			user.CartList[i].Quantity--
			if user.CartList[i].Quantity <= 0 {
				arr := []models.Item{}
				for j := 0; j < i; j++ {
					arr = append(arr, user.CartList[j])
				}
				for j := i + 1; j < l; j++ {
					arr = append(arr, user.CartList[j])
				}
				user.CartList = arr
			}
			break
		}
	}
	if !found {
		return errors.New("not found item")
	}
	if user.CartList == nil {
		user.CartList = []models.Item{}
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"cartlist": user.CartList}})
	if err != nil {
		return err
	}
	return nil

}

func AddItemOnCart(id string, item models.Item) error {
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
	found := false
	l := len(user.CartList)
	for i := 0; i < l; i++ {
		if user.CartList[i].ID == item.ID {
			qty := user.CartList[i].Quantity
			user.CartList[i].Quantity = qty + 1
			found = true
			break
		}
	}
	if !found {
		user.CartList = append(user.CartList, item)
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"cartlist": user.CartList}})
	if err != nil {
		return err
	}
	return nil
}
