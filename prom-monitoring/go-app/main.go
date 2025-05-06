package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create new serve mux
	mux := http.NewServeMux()

	// Register handlers with metrics middleware
	mux.Handle("/health", metricsMiddleware(http.HandlerFunc(healthCheck)))
	mux.Handle("/hello", metricsMiddleware(http.HandlerFunc(helloHandler)))
	mux.Handle("/data", metricsMiddleware(http.HandlerFunc(dataHandler)))

	// Add Prometheus metrics endpoint
	mux.Handle("/metrics", promhttp.Handler())

	// Start server
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", mux)
}

// Health check handler.
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Hello handler with name parameter.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}

		// Return hello message with some random status code
		statusCode := getRandomStatusCode()
		w.WriteHeader(statusCode)
		w.Write(fmt.Appendf(nil, "Hello, %s!", name))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Data handler with simulated processing delay.
func dataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		// Simulate random processing delay between 10ms to 1s
		time.Sleep(time.Duration(rand.Intn(1200)) * time.Millisecond)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": "sample data"})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getRandomStatusCode() int {
	statusCodes := []int{
		http.StatusOK,                  // 200
		http.StatusCreated,             // 201
		http.StatusAccepted,            // 202
		http.StatusNoContent,           // 204
		http.StatusMovedPermanently,    // 301
		http.StatusFound,               // 302
		http.StatusNotModified,         // 304
		http.StatusBadRequest,          // 400
		http.StatusUnauthorized,        // 401
		http.StatusForbidden,           // 403
		http.StatusNotFound,            // 404
		http.StatusTooManyRequests,     // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout,      // 504
	}
	return statusCodes[rand.Intn(len(statusCodes))]
}
