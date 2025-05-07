package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"net/http"
)

func RegisterUserRoutes() {
	http.HandleFunc("/login", controllers.LoginUser)
	http.HandleFunc("/register", controllers.RegisterationUser)
	http.Handle("/protected", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ProtectedRoutwe)))
}
