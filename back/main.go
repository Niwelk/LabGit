package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Hostname  string    `json:"hostname"`
	Timestamp string    `json:"timestamp"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		response := Response{
			Message:   "Привет из Go-бэкенда через Docker Compose!",
			Hostname:  hostname,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	})

	port := ":8080"
	println("Backend is running on port", port)
	http.ListenAndServe(port, mux)
}
