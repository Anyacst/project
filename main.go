package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.com/auth/database"
	"go.com/auth/routes"
)

func main() {

  // Initialize MySQL and Redis
	database.ConnectDBSQL()
	database.Redissetup()

	// Create router
	router := gin.Default()

	// Set up routes
	routes.InitializeRoutes(router)

	// Start the server
	port := ":8180"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}

