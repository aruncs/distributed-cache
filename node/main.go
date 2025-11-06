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

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		return
	}

	port := ":" + os.Args[1]
	// Register handlers for specific routes
	http.HandleFunc("/get", getValueHandler)
	http.HandleFunc("/put", putValueHandler)
	//http.HandleFunc("/delete", deleteValueHandler)

	// Start the HTTP server
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // Use nil for default ServeMux
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
