package main

import (
	"fmt"
	"log"
	"net/http"

	"basic-crud-ai/routes"
	"basic-crud-ai/store"
)

func main() {
	// Initialize the user store
	userStore := store.NewUserStore()

	// Add some sample data
	userStore.CreateUser("John Doe", "john@example.com")
	userStore.CreateUser("Jane Smith", "jane@example.com")

	// Setup routes
	router := routes.SetupRoutes(userStore)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Visit http://localhost:8080 for API documentation")

	log.Fatal(http.ListenAndServe(":8080", router))
}
