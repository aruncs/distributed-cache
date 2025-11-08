package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// handler for the root path "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root page!")
}

// handler for the "/hello" path
func registerNodeHandler(w http.ResponseWriter, r *http.Request) {
	var node NodeDetails

	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response := AddNode(node)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getAllNodeDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getNodeRegistry())
}

func getValueHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")

	address := getNodeAddress(key)
	url := fmt.Sprintf("%s/get?key=%s", address, key)
	resp, err := fetchURL(url)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func putValueHandler(w http.ResponseWriter, r *http.Request) {
	var data Data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	address := getNodeAddress(data.Key)
	url := fmt.Sprintf("%s/put", address)
	resp, err := postJSON(url, data.Value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func deleteValueHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Register handlers for specific routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/register-node", registerNodeHandler)
	http.HandleFunc("/nodes", getAllNodeDetails)
	http.HandleFunc("/get", getValueHandler)
	http.HandleFunc("/put", putValueHandler)
	http.HandleFunc("/delete", deleteValueHandler)

	handler := withCORS(http.DefaultServeMux)

	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handler)) // Use nil for default ServeMux
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
