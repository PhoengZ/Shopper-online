package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"net/http"
)

func RegisterUserRoutes() {
	http.HandleFunc("/auth/login", controllers.LoginUser)
	http.HandleFunc("/auth/register", controllers.RegisterationUser)
	http.Handle("/profile", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ProtectedRoute)))
}
