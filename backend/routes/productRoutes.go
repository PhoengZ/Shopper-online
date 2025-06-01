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
	r.Handle("/auth/getStoreItem/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetProductSellingByID))).Methods("GET")
	r.Handle("/auth/addStoreItem", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateItem))).Methods("POST")
	r.Handle("/auth/editStoreItem", middlewares.AuthMiddleware(http.HandlerFunc(controllers.EditItem))).Methods("PATCH")
	r.Handle("/auth/removeStoreItem", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeleteItem))).Methods("DELETE")
}
