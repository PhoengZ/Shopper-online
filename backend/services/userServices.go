package services

import (
	"context"
	"errors"
	"time"

	"backend/config"
	"backend/models"
	"backend/utils"

	"go.mongodb.org/mongo-driver/bson"
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
	token, err := utils.GenerateToken(user.ID)
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
