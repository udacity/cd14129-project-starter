package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":8080", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func validateUser(user User) error {
	// TODO: Implement validation logic
	return nil
}
