package main

// The data model
type User struct {
	ID    string `json:"id"` // Tag for JSON serialization
	Name  string `json:"name"`
	Email string `json:"email"`
}

// sample database
var users = map[string]User{
	"1": {ID: "1", Name: "John Doe", Email: "john@example.com"},
	"2": {ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
}

//Method to access data on sample db
//Handler function
