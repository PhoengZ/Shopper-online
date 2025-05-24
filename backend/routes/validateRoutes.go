package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

func RegisterValidateRoutes(r *mux.Router) {
	r.HandleFunc("/validate", controllers.ValidToken).Methods("POST")
}
