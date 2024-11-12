package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/radenrishwan/monitoring/metrics"
)

func initLogger() {
	err := os.MkdirAll("/app/logs", 0755)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile("/app/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	slog.SetDefault(logger)
}

func generateRequestLog(path string, method string, message string) {
	slog.Info("Request incomming", "path", path, "method", method, "message", message)
}

func main() {
	initLogger()

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		generateRequestLog(r.URL.Path, r.Method, "")

		w.Write([]byte("Hello, World!"))

		// TODO: move into middleware later
		metrics.RequestCounter.WithLabelValues(r.URL.Path, r.Method, http.StatusText(http.StatusOK)).Inc()
		metrics.RequestDuration.
			WithLabelValues(r.URL.Path, r.Method, http.StatusText(http.StatusOK)).
			Observe(time.Since(start).Seconds())
	})

	server.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		generateRequestLog(r.URL.Path, r.Method, "")

		w.Write([]byte("pong"))

		metrics.RequestCounter.WithLabelValues(r.URL.Path, r.Method, http.StatusText(http.StatusOK)).Inc()
		metrics.RequestDuration.
			WithLabelValues(r.URL.Path, r.Method, http.StatusText(http.StatusOK)).
			Observe(time.Since(start).Seconds())
	})

	server.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", server); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
