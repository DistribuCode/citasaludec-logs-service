package config

import (
    "log"
    "github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️ No se encontró .env, se usarán variables de entorno del sistema")
    }
}
