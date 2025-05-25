package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils/response"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("userID")
	response.JSONResponse(w, http.StatusOK, map[string]string{"message": "Hello, " + userID})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	token, err := services.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		response.JSONResponse(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]string{"token": token})
}
func RegisterationUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
		return
	}
	user.CartList = []models.Item{}
	err = services.RegisterUser(user)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}

	response.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func GetCartList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]
	cartList, err := services.GetCartListByID(Id)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"message": err.Error()})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]interface{}{"products": cartList})
}
func AddItemToCart(w http.ResponseWriter, r *http.Request) {
	var object struct {
		Id      string      `json:"id"`
		Product models.Item `json:"product"`
	}
	err := json.NewDecoder(r.Body).Decode(&object)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	err = services.AddItemOnCart(object.Id, object.Product)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"message": err.Error()})
	}
	response.JSONResponse(w, http.StatusOK, map[string]string{"message": "Success adding item"})
}
