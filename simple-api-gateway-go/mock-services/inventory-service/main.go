package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("INVENTORY_SERVICE_PORT", "8083")
	mux := http.NewServeMux()

	mux.HandleFunc("/api/inventory/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"service": "inventory-service",
			"items": []map[string]interface{}{
				{"sku": "PRD-001", "name": "Duff Beer", "stock": 32},
				{"sku": "PRD-002", "name": "Donut Box", "stock": 12},
			},
		})
	})

	mux.HandleFunc("/api/inventory/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"service": "inventory-service",
			"status":  "ok",
		})
	})

	addr := ":" + port
	log.Printf("inventory-service running on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("inventory-service failed: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
