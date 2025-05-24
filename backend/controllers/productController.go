package controllers

import (
	"backend/services"
	"backend/utils/response"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	listProduct, err := services.GetProducts()
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
		return
	}
	response.JSONResponse(w, http.StatusOK, listProduct)
}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	ProductID := r.URL.Query().Get("id")
	product, err := services.GetProductByID(ProductID)
	if err != nil {
		if err.Error() == "not_found" {
			response.JSONResponse(w, http.StatusNotFound, map[string]string{
				"error": "Product not found",
			})
			return
		}
		response.JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve product",
		})
		return
	}
	response.JSONResponse(w, http.StatusOK, product)
}
