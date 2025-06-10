package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterationTopUp(r *mux.Router) {
	r.Handle("/topup/request", middlewares.AuthMiddleware(http.HandlerFunc(controllers.Request))).Methods("POST")
	r.Handle("/topup/history", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetHistory))).Methods("GET")
}
