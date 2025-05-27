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
	r.Handle("/auth/getProfile/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
	r.Handle("/auth/updateProfile", middlewares.AuthMiddleware(http.HandlerFunc(controllers.UpdateProfile))).Methods("PATCH")
	r.Handle("/addCartItem", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AddItemToCart))).Methods("POST")
	r.Handle("/removeCartItem", middlewares.AuthMiddleware(http.HandlerFunc(controllers.RemoveItemOnCart))).Methods("PATCH")
	r.Handle("/cartList/{id}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetCartList))).Methods("GET")
}
