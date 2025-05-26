package controllers

import (
	"backend/utils/response"
	"encoding/json"
	"net/http"
	"os"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data/categories.json")
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to read categories data"})
		return
	}
	var categories []string
	err = json.Unmarshal(data, &categories)
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": " Failed to parse categories data"})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]interface{}{"categories": categories})
}
