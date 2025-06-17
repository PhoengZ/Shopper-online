package controllers

import (
	"backend/utils/response"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	var path string
	var data []byte
	var err error
	if os.Getenv("VERCEL") == "" {
		path = "public/data/categories.json"
		data, err = os.ReadFile(path)
		if err != nil {
			response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to read categories data"})
			return
		}
	} else {
		path = "https://shopper-online.vercel.app/data/categories.json"
		resp, err := http.Get(path)
		if err != nil {
			response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch categories data"})
			return
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to read categories data"})
			return
		}
	}
	var categories []string
	err = json.Unmarshal(data, &categories)
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": " Failed to parse categories data"})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]interface{}{"categories": categories})
}
