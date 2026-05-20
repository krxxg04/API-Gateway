package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("PAYMENT_SERVICE_PORT", "8084")
	mux := http.NewServeMux()

	mux.HandleFunc("/api/payments/charge", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"service": "payment-service",
			"result":  "payment authorized",
			"amount":  99.90,
		})
	})

	mux.HandleFunc("/api/payments/history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"service": "payment-service",
			"history": []map[string]interface{}{
				{"id": "PMT-1001", "amount": 49.90, "status": "approved"},
				{"id": "PMT-1002", "amount": 99.90, "status": "approved"},
			},
		})
	})

	addr := ":" + port
	log.Printf("payment-service running on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("payment-service failed: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
