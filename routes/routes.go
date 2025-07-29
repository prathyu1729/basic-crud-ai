package routes

import (
	"github.com/gorilla/mux"

	"basic-crud-ai/handlers"
	"basic-crud-ai/store"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(userStore *store.UserStore) *mux.Router {
	r := mux.NewRouter()

	// Create handlers
	userHandler := handlers.NewUserHandler(userStore)
	homeHandler := handlers.NewHomeHandler()

	// Routes
	r.HandleFunc("/", homeHandler.ServeHome).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return r
}
