package controllers

import (
	"backend/utils/response"
	"net/http"
)

func Request(w http.ResponseWriter, r *http.Request) {
	response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "test failed"})
}

func GetHistory(w http.ResponseWriter, r *http.Request) {

}
