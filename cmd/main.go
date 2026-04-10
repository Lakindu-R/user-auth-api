package main

import (
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

	fmt.Println("Server running on port " + port + "...")
	http.ListenAndServe(":"+port, nil)
}
