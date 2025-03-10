package main

import (
	"log"
	"net/http"
	"rate-limiter/config"
	"rate-limiter/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Carregar configurações
	config.LoadEnv()

	// Configurar roteador
	r := mux.NewRouter()

	// Adicionar middleware de rate limiter
	r.Use(middleware.RateLimiterMiddleware)

	// Rota de exemplo
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Rate Limiter!"))
	}).Methods("GET")

	// Iniciar servidor
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
