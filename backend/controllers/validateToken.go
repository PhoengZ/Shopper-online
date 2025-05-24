package controllers

import (
	"backend/utils"
	"backend/utils/response"
	"encoding/json"
	"net/http"
)

func ValidToken(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Failed to parse JSON"})
		return
	}

	userID, err := utils.ValidateToken(credentials.Token)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid token"})
		return
	}

	response.JSONResponse(w, http.StatusOK, map[string]string{
		"userID":  userID,
		"message": "Valid",
	})
}
