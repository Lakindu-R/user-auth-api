package routes

import (
	"net/http"
	"user-auth-api/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/protected", controllers.Protected)
}