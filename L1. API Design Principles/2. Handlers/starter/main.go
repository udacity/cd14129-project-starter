package main

import (
	"log"
	"net/http"
)

func main() {
	// TODO: Register your handlers here

	// TODO: Serve on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO: Define greet and userInfo handler functions here
