package main

import (
    "logs-service/config"
    "logs-service/routes"
    "logs-service/services"  // 👈 importa el paquete donde está tu StartRabbitConsumer
    "log"
    "net/http"
    "os"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    config.LoadEnv()

    // 👉 Iniciar el consumidor RabbitMQ en un goroutine
    go services.StartRabbitConsumer()

    r := routes.Setup()

    // Exponer endpoint Prometheus
    r.Handle("/metrics", promhttp.Handler())

    port := os.Getenv("PORT")
    log.Printf("🚀 Logs Service escuchando en puerto :%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
