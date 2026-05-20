package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("USER_SERVICE_PORT", "8082")
	mux := http.NewServeMux()

	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"service": "user-service",
			"users": []map[string]interface{}{
				{"id": 1, "name": "Lisa"},
				{"id": 2, "name": "Milhouse"},
			},
		})
	})

	mux.HandleFunc("/api/users/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"service": "user-service",
			"profile": map[string]interface{}{
				"id":    1,
				"name":  "Lisa",
				"email": "lisa@example.com",
			},
		})
	})

	addr := ":" + port
	log.Printf("user-service running on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("user-service failed: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
