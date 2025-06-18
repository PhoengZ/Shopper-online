package controllers

import (
	"backend/utils/response"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	project_url := os.Getenv("SUPABASE_URL")
	bucketName := "json"
	filepath := "categories.json"
	api_key := os.Getenv("API_KEY")
	download_url := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", project_url, bucketName, filepath)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", download_url, nil)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"error": "Failed to create request"})
		return
	}
	req.Header.Set("apikey", api_key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"error": "Failed to fetch data"})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"error": "Get categories failed"})
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"error": "Failed parse data"})
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
