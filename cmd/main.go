package main

import (
    "logs-service/config"
    "logs-service/routes"
    "log"
    "net/http"
    "os"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    config.LoadEnv()

    r := routes.Setup()

    // Exponer endpoint Prometheus
    r.Handle("/metrics", promhttp.Handler())

    port := os.Getenv("PORT")
    log.Printf("🚀 Logs Service escuchando en puerto :%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
