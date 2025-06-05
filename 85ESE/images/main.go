package main

import (
	"images/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
)

func main() {
	// Configurar logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "product-images",
		Level: hclog.LevelFromString("DEBUG"),
	})

	// Criar armazenamento (exemplo fictício)
	store := handlers.NewLocalStorage("./imagestore", logger)

	// Criar handler
	fh := handlers.NewFiles(store, logger)

	// Criar roteador
	router := mux.NewRouter()

	// Registrar rotas
	router.HandleFunc("/upload", fh.UploadMultipart).Methods(http.MethodPost)
	router.HandleFunc("/ping", fh.Ping).Methods(http.MethodGet)
	router.HandleFunc("/images/{id}", fh.ServeProductImage).Methods("GET")

	// Iniciar servidor
	port := ":9091"
	log.Printf("Iniciando servidor na porta %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
