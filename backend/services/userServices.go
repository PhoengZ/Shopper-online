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

func GetProfile(id string) (map[string]interface{}, error) {
	collection := config.GetCollection("User")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return map[string]interface{}{}, err
	}
	profile := map[string]interface{}{
		"username": user.Username,
		"address":  user.Address,
		"coin":     user.Coin,
		"history":  user.History,
	}
	return profile, nil
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

func UpdateProfile(id string, newProfile map[string]interface{}) error {
	collection := config.GetCollection("User")
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID format")
	}
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}
	updateFields := bson.M{}
	if newProfile["password"] != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newProfile["password"].(string)), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash password")
		}
		updateFields["password"] = string(hashedPassword)
	}
	if newProfile["address"] != "" {
		updateFields["address"] = newProfile["address"]
	}
	if newProfile["coin"] != 0 {
		newCoin := newProfile["coin"].(float64)
		intCoin := int(newCoin)
		updateFields["coin"] = intCoin + user.Coin
	}
	if newProfile["history"] != nil {
		newHistory := user.History
		newItem := newProfile["history"].(map[string]interface{})
		if len(newItem) > 0 {
			newHistory = append([]models.Item{newProfile["history"].(models.Item)}, newHistory...)
			updateFields["history"] = newHistory
		}
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateFields})
	if err != nil {
		return errors.New("failed to update profile")
	}
	return nil
}
