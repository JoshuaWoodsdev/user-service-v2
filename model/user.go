package model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users = map[string]User{
	"1": {
		ID:    "1",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	},
	"2": {
		ID:    "2",
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	},
	"3": {
		ID:    "3",
		Name:  "Alice Smith",
		Email: "alice.smith@example.com",
	},
