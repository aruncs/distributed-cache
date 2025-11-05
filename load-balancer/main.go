package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler for the root path "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root page!")
}

// handler for the "/hello" path
func registerNodeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Register handlers for specific routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register-node",registerNodeHandler)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // Use nil for default ServeMux
}