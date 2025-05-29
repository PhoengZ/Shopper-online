package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.GetProductByID).Methods("GET")
	r.Handle("/auth/getSelling/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetProductSellingByID))).Methods("GET")
	r.Handle("/auth/createItem/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateItem))).Methods("POST")
	r.Handle("/auth/editItem/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.EditItem))).Methods("PATCH")
}
