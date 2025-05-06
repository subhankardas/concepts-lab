package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"path", "method", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: []float64{0.1, 0.3, 0.5, 1, 1.5, 2, 2.5, 5, 7.5, 10},
	}, []string{"path", "method"})
)

// metricsMiddleware is an HTTP middleware that records the duration and count of HTTP requests.
func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &StatusRecorder{ResponseWriter: w, Status: 200}

		// Call the next handler
		next.ServeHTTP(recorder, r)

		duration := time.Since(start).Seconds()
		httpRequestDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration)
		httpRequestsTotal.WithLabelValues(r.URL.Path, r.Method, fmt.Sprintf("%d", recorder.Status)).Inc()
	})
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
