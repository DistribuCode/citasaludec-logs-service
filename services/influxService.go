package services

import (
	"context"
    "log"
    "os"
    "time"

	influxdb2 "github.com/influxdata/influxdb-client-go"


)

func SaveLogToInflux(evento, destinatario, estado, canal, errorMsg string, timestamp time.Time) {
    client := influxdb2.NewClient(os.Getenv("INFLUX_URL"), os.Getenv("INFLUX_TOKEN"))
    defer client.Close()

    writeAPI := client.WriteApiBlocking(os.Getenv("INFLUX_ORG"), os.Getenv("INFLUX_BUCKET"))

    p := influxdb2.NewPoint("logs_entrega",
        map[string]string{
            "evento": evento,
            "canal":  canal,
        },
        map[string]interface{}{
            "destinatario": destinatario,
            "estado":       estado,
            "error":        errorMsg,
        },
        timestamp,
    )

    err := writeAPI.WritePoint(context.Background(), p)
    if err != nil {
        log.Printf("❌ Error escribiendo en InfluxDB: %v\n", err)
    } else {
        log.Println("✅ Log almacenado en InfluxDB")
    }
}
