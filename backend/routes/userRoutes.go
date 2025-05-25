package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/auth/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/auth/register", controllers.RegisterationUser).Methods("POST")
	r.HandleFunc("/cartList/{id}", controllers.GetCartList).Methods("GET")
	r.HandleFunc("/addCartItem", controllers.AddItemToCart).Methods("POST")
	r.Handle("/profile", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ProtectedRoute))).Methods("GET")
}
