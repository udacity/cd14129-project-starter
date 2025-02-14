package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// Set up the route for the register endpoint and start the server on port 8080
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":8080", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure that only POST requests are allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	// Decode the JSON body of the request into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// If the JSON is malformed or cannot be decoded, return a 400 error
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Validate the user input data
	err = validateUser(user)
	if err != nil {
		// If validation fails, return a 400 error with the specific error message
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If the input is valid, respond with a 200 status and success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Registration successful!")
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func validateUser(user User) error {
	// Check if the name is empty
	if user.Name == "" {
		return errors.New("Name cannot be empty")
	}

	// Check if the email contains an "@" symbol
	if !strings.Contains(user.Email, "@") {
		return errors.New("Invalid email address")
	}

	// Ensure that the age is greater than 18
	if user.Age <= 18 {
		return errors.New("Age must be greater than 18")
	}

	return nil
}
