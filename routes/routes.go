package routes

import (
    "logs-service/controllers"
    "logs-service/middlewares"
    "github.com/gorilla/mux"
)

func Setup() *mux.Router {
    r := mux.NewRouter()
    r.Use(middlewares.AuthMiddleware)

    r.HandleFunc("/logs", controllers.ReceiveLog).Methods("POST")

    return r
}
