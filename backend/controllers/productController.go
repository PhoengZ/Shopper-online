package controllers

import (
	"backend/services"
	"backend/utils/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("name")
	categoryParam := r.URL.Query().Get("category")
	var category []string
	if categoryParam == "" {
		category = []string{}
	} else {
		category = strings.Split(categoryParam, ",")
	}
	Sprice := r.URL.Query().Get("price")
	Iprice, err := strconv.ParseBool(Sprice)
	if err != nil {
		Iprice = true
	}
	listProduct, err := services.GetProducts(itemName, category, Iprice)
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
		return
	}
	response.JSONResponse(w, http.StatusOK, listProduct)
}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProductID := vars["id"]
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

func GetProductSellingByID(w http.ResponseWriter, r *http.Request) {

}
func CreateItem(w http.ResponseWriter, r *http.Request) {

}

func EditItem(w http.ResponseWriter, r *http.Request) {

}
