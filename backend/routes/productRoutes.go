package routes

import (
	"backend/controllers"
	"net/http"
)

func RegisterProductRoutes() {
	http.HandleFunc("/products", controllers.GetProducts)
	http.HandleFunc("/products/{id}", controllers.GetProductByID)
}
