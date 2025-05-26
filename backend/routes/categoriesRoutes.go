package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterCategoriesRoutes(r *mux.Router) {
	r.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
}
