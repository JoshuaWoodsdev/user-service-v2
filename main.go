package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// 1. Import the Gorilla Mux package
	"github.com/gorilla/mux"
)

// The data model
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Sample database
var users = map[string]User{
	"1": {ID: "1", Name: "John Doe", Email: "john@example.com"},
	"2": {ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
}

// --- Handler Function (Using Mux) ---
func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	// 1. Use mux.Vars() to automatically extract variables defined in the route pattern.
	vars := mux.Vars(r)
	userID := vars["id"] // We extract the variable named 'id'

	// 2. Data Retrieval (The business logic remains the same)
	user, ok := users[userID]
	if !ok {
		// Log and return HTTP 404
		log.Printf("User ID %s not found in map", userID)
		http.Error(w, fmt.Sprintf("User ID %s not found", userID), http.StatusNotFound)
		return
	}

	// 3. Response Generation
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		// Log and return HTTP 500
		log.Printf("Error encoding response for user %s: %v", userID, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Implicitly returns HTTP 200 OK
}

// --- Server Initialization (Using Mux Router) ---
func main() {
	// 1. Create a new Mux router instance
	router := mux.NewRouter()

	// 2. Define the route pattern with a variable: {id}
	router.HandleFunc("/api/users/{id}", GetUserHandler).Methods("GET")

	// Health check (optional, but good practice)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Service is healthy! (Gorilla Mux)")
	}).Methods("GET")

	port := ":8080"
	log.Printf("🚀 Starting User Microservice (Mux Version) on http://localhost%s", port)

	// 3. Start the server, passing the Mux router as the handler
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
