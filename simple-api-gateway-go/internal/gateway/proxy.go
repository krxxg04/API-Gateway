package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"simple-api-gateway-go/internal/response"
)

func NewReverseProxy(target string) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)
		r.Header.Set("X-Forwarded-Host", r.Host)
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		response.WriteError(w, http.StatusBadGateway, "upstream service unavailable")
	}

	return proxy, nil
}
