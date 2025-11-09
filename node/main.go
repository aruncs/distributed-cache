package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Key struct {
	Key string `json:"key"`
}

var isOOR bool = false

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		return
	}

	handler := withCORS(http.DefaultServeMux)

	port := ":" + os.Args[1]
	// Register handlers for specific routes
	http.HandleFunc("/health-check", getHealthCheckHandler)
	http.HandleFunc("/get", getValueHandler)
	http.HandleFunc("/put", putValueHandler)
	http.HandleFunc("/delete", deleteValueHandler)
	http.HandleFunc("/data", getDataHandler)

	// Start the HTTP server
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handler)) // Use nil for default ServeMux
}

func getValueHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	value := getValue(key)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)

}

func putValueHandler(w http.ResponseWriter, r *http.Request) {
	var data Data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	putValue(data.Key, data.Value)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ok")

}

func deleteValueHandler(w http.ResponseWriter, r *http.Request) {
	var data Key

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	deleteValue(data.Key)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ok")
}

func getHealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	if isOOR {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode("not ok")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ok")
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	store := getStore()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // or specific origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
