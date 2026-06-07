package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
	Database  string `json:"database"`
}

func connectMongoDB() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Println("MONGO_URI not set, skipping MongoDB connection")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Printf("MongoDB connection failed: %v", err)
		return
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("MongoDB ping failed: %v", err)
		return
	}

	mongoClient = client
	log.Println("MongoDB connected successfully")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	dbStatus := "disconnected"
	if mongoClient != nil {
		dbStatus = "connected"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "1.0.0",
		Database:  dbStatus,
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "StartTech API is running",
	})
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("method=%s path=%s duration=%s",
			r.Method, r.URL.Path, time.Since(start))
	}
}

func main() {
	connectMongoDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	http.HandleFunc("/", loggingMiddleware(homeHandler))

	fmt.Printf("StartTech backend starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
