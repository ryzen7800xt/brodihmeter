package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CounterData struct {
	Meanness int `json:"meanness"`
}

type HistoryEntry struct {
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}

type HistoryResponse struct {
	Entries []HistoryEntry `json:"entries"`
}

var (
	counter   = 0
	history   = []HistoryEntry{}
	counterMu sync.Mutex
)

func main() {
	fmt.Println("Brodih Program")
	fmt.Println("Starting server on port 8080...")
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
			history = append(history, HistoryEntry{
				Count:     counter,
				Timestamp: time.Now(),
			})
			defer counterMu.Unlock()
			json.NewEncoder(w).Encode(CounterData{Meanness: counter})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// API endpoint to reset counter
	http.HandleFunc("/api/reset", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			counterMu.Lock()
			counter = 0
			history = []HistoryEntry{}
			defer counterMu.Unlock()
			json.NewEncoder(w).Encode(CounterData{Meanness: 0})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// API endpoint to get history
	http.HandleFunc("/api/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		counterMu.Lock()
		defer counterMu.Unlock()
		json.NewEncoder(w).Encode(HistoryResponse{Entries: history})
	})

	http.ListenAndServe(":8080", nil)
}
