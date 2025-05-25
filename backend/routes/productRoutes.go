package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.GetProductByID).Methods("GET")
}
