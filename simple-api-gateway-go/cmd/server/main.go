package main

import (
	"log"
	"net/http"

	"simple-api-gateway-go/internal/config"
	"simple-api-gateway-go/internal/gateway"
	"simple-api-gateway-go/internal/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router, err := gateway.NewRouter(cfg)
	if err != nil {
		log.Fatalf("failed to build router: %v", err)
	}

	handler := middleware.RecoveryMiddleware(middleware.LoggingMiddleware(router))

	addr := ":" + cfg.Port
	log.Printf("API Gateway listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
