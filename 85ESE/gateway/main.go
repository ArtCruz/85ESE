package main

import (
	"gateway/config"
	"gateway/internal/handlers"
	"log"
	"net/http"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Carregar configurações
	cfg := config.Load()

	// Criar roteador
	router := mux.NewRouter()

	// Servir arquivos estáticos da pasta "static"
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/").Handler(fs)

	// Registrar handlers
	handlers.RegisterRoutes(router, cfg)

	// Configurar CORS
	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:7071"}), // Origem do frontend
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		gohandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      ch(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Gateway rodando na porta %s", cfg.ServerAddress)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
