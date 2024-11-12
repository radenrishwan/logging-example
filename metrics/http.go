package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var RequestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "The total number of processed events",
}, []string{"path", "method", "status"})

var RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_request_duration_seconds",
	Help: "The duration of the request",
}, []string{"path", "method", "status"})
