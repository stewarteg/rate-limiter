package main

import (
	"log"
	"net/http"
	"rate-limiter/config"
	"rate-limiter/middleware"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	r := mux.NewRouter()

	// Adicionar middleware de rate limiter
	r.Use(middleware.RateLimiterMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Rate Limiter!"))
	}).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
