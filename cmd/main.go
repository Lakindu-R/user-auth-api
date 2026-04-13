package main
// @title User Auth API
// @version 1.0
// @description This is a JWT authentication API in Go
// @host localhost:8081
// @BasePath /
import (
	_ "user-auth-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"fmt"
	"net/http"
	"os"

	"user-auth-api/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	routes.SetupRoutes()
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("Server running on port " + port + "...")
	http.ListenAndServe(":"+port, nil)
}
