package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"backend/utils/response"
	"encoding/json"
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
	vars := mux.Vars(r)
	uid := vars["id"]
	item, err := services.GetSellingByID(uid)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"message": "Failed to get item"})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]interface{}{"products": item})
}
func CreateItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB litmit
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Failed to parse form data because to large"})
		return
	}
	imageURL, err := utils.UploadImage("file", r)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"message": err.Error()})
		return
	}
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Invalid price format"})
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Invalid quantity format"})
		return
	}
	ctg := strings.Split(r.FormValue("category"), ",")
	if len(ctg) == 0 || (len(ctg) == 1 && ctg[0] == "") {
		ctg = []string{}
	}
	item := models.Product{
		UserID:      r.FormValue("uid"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Quantity:    quantity,
		Category:    ctg,
		URL:         imageURL,
	}
	product, err := services.CreateItem(item)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"message": err.Error()})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]interface{}{"product": product})
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Failed to parse form data because to large"})
		return
	}
	check, err := utils.HavingFieldImage(r)
	if err != nil {
		if check {
			response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Failed to parse form data because something"})
			return
		}
	}
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Invalid price format"})
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Invalid quantity format"})
		return
	}
	ctg := strings.Split(r.FormValue("category"), ",")
	if len(ctg) == 0 || (len(ctg) == 1 && ctg[0] == "") {
		ctg = []string{}
	}
	item := models.Product{
		ID:          r.FormValue("id"),
		UserID:      r.FormValue("uid"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Price:       price,
		Quantity:    quantity,
		Category:    ctg,
	}
	if check {
		item.URL, err = utils.UploadImage("file", r)
		if err != nil {
			response.JSONResponse(w, http.StatusConflict, map[string]string{"message": "failed to upload image"})
			return
		}
	} else {
		item.URL = ""
	}
	product, err := services.EditItem(item)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"message": err.Error()})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]interface{}{"product": product})
}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		PID string `json:"productID"`
	}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
		return
	}
	err = services.DeleteItem(payload.PID)
	if err != nil {
		response.JSONResponse(w, http.StatusConflict, map[string]string{"message": "Failed to delete product"})
		return
	}
	response.JSONResponse(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
