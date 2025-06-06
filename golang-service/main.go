package main

import (
	"fmt"
	"golang-service/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.HandleFunc("/go-service/health", handlers.HealthCheck)
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting server on port %s", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting server: ")
	}
}
