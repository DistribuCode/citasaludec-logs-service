package controllers

import (
    "encoding/json"
    "logs-service/services"
    "net/http"
    "time"
)

type LogRequest struct {
    Evento      string `json:"evento"`
    Destinatario string `json:"destinatario"`
    Estado      string `json:"estado"`
    Canal       string `json:"canal"`
    ErrorMsg    string `json:"error"`
}

func ReceiveLog(w http.ResponseWriter, r *http.Request) {
    var req LogRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    now := time.Now()

    // Guardar en Influx
    services.SaveLogToInflux(req.Evento, req.Destinatario, req.Estado, req.Canal, req.ErrorMsg, now)

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"Log recibido y almacenado"}`))
}
