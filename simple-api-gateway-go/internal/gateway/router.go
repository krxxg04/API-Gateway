package gateway

import (
	"encoding/json"
	"net/http"

	"simple-api-gateway-go/internal/config"
	"simple-api-gateway-go/internal/middleware"
)

func NewRouter(cfg *config.Config) (http.Handler, error) {
	authProxy, err := NewReverseProxy(cfg.AuthServiceURL)
	if err != nil {
		return nil, err
	}

	userProxy, err := NewReverseProxy(cfg.UserServiceURL)
	if err != nil {
		return nil, err
	}

	inventoryProxy, err := NewReverseProxy(cfg.InventoryServiceURL)
	if err != nil {
		return nil, err
	}

	paymentProxy, err := NewReverseProxy(cfg.PaymentServiceURL)
	if err != nil {
		return nil, err
	}

	jwtMiddleware := middleware.JWTAuthMiddleware(cfg.JWTSecret)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
			"service": "simple-api-gateway-go",
		})
	})

	// Public routes
	mux.Handle("/api/auth", authProxy)
	mux.Handle("/api/auth/", authProxy)

	// Protected routes
	mux.Handle("/api/users", jwtMiddleware(userProxy))
	mux.Handle("/api/users/", jwtMiddleware(userProxy))

	mux.Handle("/api/inventory", jwtMiddleware(inventoryProxy))
	mux.Handle("/api/inventory/", jwtMiddleware(inventoryProxy))

	mux.Handle("/api/payments", jwtMiddleware(paymentProxy))
	mux.Handle("/api/payments/", jwtMiddleware(paymentProxy))

	return mux, nil
}
