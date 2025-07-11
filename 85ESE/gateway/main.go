package main

import (
	"gateway/config"
	"gateway/internal/handlers"
	"log"
	"net/http"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"

)

func main() {
	// Carregar configurações
	cfg := config.Load()

	// Criar roteador
	router := mux.NewRouter()

	// ➕ Adicionar rota para métricas do Prometheus
	router.Handle("/metrics", promhttp.Handler())


	// REGISTRE AS ROTAS PRIMEIRO!
	handlers.RegisterRoutes(router, cfg)

	// Depois, sirva arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/").Handler(fs)

	// Configurar CORS
	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:7071"}),
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
