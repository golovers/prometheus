package service

import (
	"github.com/felixge/httpsnoop"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapi_http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
		},
		[]string{"code", "method", "path"},
	)
	latencyHistogram = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "myapi_http_requests_duration_millisecond",
			Help: "How long it took to process the request, partitioned by status code, method and HTTP path.",
		},
		[]string{"code", "method", "path"},
	)
	inFlightGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "myapi_http_requests_active",
			Help: "How many active in-flight requests, partitioned by status code, method, HTTP path.",
		},
		[]string{"method", "path"},
	)
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inFlightGauge.WithLabelValues(r.Method, r.URL.Path).Inc()
		metrics := httpsnoop.CaptureMetrics(next, w, r)

		requestCounter.WithLabelValues(http.StatusText(metrics.Code), r.Method, r.URL.Path).Inc()

		inFlightGauge.WithLabelValues(r.Method, r.URL.Path).Dec()
		latencyHistogram.WithLabelValues(http.StatusText(metrics.Code), r.Method, r.URL.Path).Observe(float64(metrics.Duration / time.Millisecond))

	})
}

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(latencyHistogram)
	prometheus.MustRegister(inFlightGauge)
}
