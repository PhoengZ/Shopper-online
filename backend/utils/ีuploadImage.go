package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func UploadImage(fieldName string, r *http.Request) (string, error) {
	file, handler, err := r.FormFile(fieldName)
	if err != nil {
		if err == http.ErrMissingFile {
			return "https://dubccshnzuayvwaiudox.supabase.co/storage/v1/object/public/json/productImage/images.png", nil
		}
		return "", err
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filepath := "productImage/" + handler.Filename
	apikey := os.Getenv("API_KEY")
	supabase_url := os.Getenv("SUPABASE_URL")
	bucket := "json"
	url, err := UploadFileToSupabase(ctx, bucket, filepath, apikey, supabase_url, file)
	if err != nil {
		return "", err
	}

	return url, nil
}

func HavingFieldImage(r *http.Request) (bool, error) {
	_, _, err := r.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			return false, err
		}
		return true, err
	}
	return true, nil
}

func UploadFileToSupabase(ctx context.Context, bucket, filepath, apikey, supabase_url string, file io.Reader) (string, error) {
	uploadUrl := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabase_url, bucket, filepath)
	req, err := http.NewRequestWithContext(ctx, "POST", uploadUrl, file)
	if err != nil {
		return "", err
	}
	req.Header.Set("apikey", apikey)
	req.Header.Set("Authorization", "Bearer "+apikey)
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("upload failed: status=%d, response=%s", res.StatusCode, string(body))
	}
	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabase_url, bucket, filepath)
	return publicURL, nil
}
