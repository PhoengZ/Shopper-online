package controllers

import (
	"backend/services"
	"backend/utils/response"
	"encoding/json"
	"net/http"
)

func Request(w http.ResponseWriter, r *http.Request) {
	var object struct {
		UserID string `json:"userID"`
		Amount int    `json:"amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&object)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Bad payload"})
		return
	}
	err = services.MakeRequest(object.UserID, object.Amount)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]string{"message": "Success requset"})
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	transaction, err := services.GetHistory()
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	response.JSONResponse(w, http.StatusOK, transaction)
}
