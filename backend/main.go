package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "1.0.0",
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ApiResponse{
		Message: "StartTech API is running",
	})
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf(
			"method=%s path=%s duration=%s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", loggingMiddleware(healthHandler))
	mux.HandleFunc("/", loggingMiddleware(homeHandler))

	fmt.Printf("StartTech backend starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
