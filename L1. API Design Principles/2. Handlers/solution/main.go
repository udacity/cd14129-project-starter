package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Register handlers for the routes
	http.HandleFunc("/", greet)
	http.HandleFunc("/user", userInfo)

	// Serve on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// greet responds with a simple greeting message
func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my server!"))
}

// userInfo responds with a mock user data in JSON format
func userInfo(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
