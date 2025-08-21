package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	appPort := getEnv("APP_PORT", "8080")
	metricsPort := getEnv("METRICS_PORT", "8081")

	// ... код для запуска серверов ...

	go func() {
		addr := fmt.Sprintf(":%s", appPort)
		fmt.Printf("Starting main server on %s\n", addr)
		if err := http.ListenAndServe(addr, mainMux); err != nil {
			// ...
		}
	}()
	
	// ...

	addr := fmt.Sprintf(":%s", metricsPort)
	fmt.Printf("Starting metrics server on %s\n", addr)
	if err := http.ListenAndServe(addr, metricsMux); err != nil {
		// ...
	}
}