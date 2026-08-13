package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type CounterData struct {
	Meanness int `json:"meanness"`
}

var (
	counter   = 0
	counterMu sync.Mutex
)

func main() {
	fmt.Println("Brodih Program")
	fmt.Println("Starting server on port 8080...")
 // so stupid such gay code
	// Serve static files from current directory
	http.Handle("/", http.FileServer(http.Dir(".")))

	// API endpoint to get current count
	http.HandleFunc("/api/count", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		counterMu.Lock()
		defer counterMu.Unlock()
		json.NewEncoder(w).Encode(CounterData{Meanness: counter})
	})

	// API endpoint to increment count
	http.HandleFunc("/api/increment", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			counterMu.Lock()
			counter++
			defer counterMu.Unlock()
			json.NewEncoder(w).Encode(CounterData{Meanness: counter})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
