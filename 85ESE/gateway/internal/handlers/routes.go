package handlers

import (
	"fmt"
	"gateway/config"
	"gateway/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, cfg *config.Config) {
	// Rotas para o serviço images
	router.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		// Log da requisição recebida do frontend
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)

		// Encaminhar a requisição para o serviço images
		err := services.UploadImage(cfg.ImagesAPIURL, r)
		if err != nil {
			fmt.Printf("Erro ao enviar imagem para o serviço images: %v\n", err)
			http.Error(w, "Erro ao enviar imagem", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Imagem enviada com sucesso"))
	}).Methods(http.MethodPost)

	router.HandleFunc("/ping/images", func(w http.ResponseWriter, r *http.Request) {
		// URL do serviço images
		url := cfg.ImagesAPIURL + "/ping"

		// Enviar requisição GET para o serviço images
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Erro ao pingar o serviço images: %v\n", err)
			http.Error(w, "Serviço images não está acessível", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Log da resposta do serviço images
		fmt.Printf("Resposta do serviço images: %d %s\n", resp.StatusCode, resp.Status)

		// Retornar a resposta para o cliente
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(fmt.Sprintf("Serviço images respondeu com status: %s", resp.Status)))
	}).Methods(http.MethodGet)

	// Rota para o serviço products
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		// Log da requisição recebida do frontend
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)

		products, err := services.FetchProducts(cfg.ProductAPIURL)
		if err != nil {
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(products)
	}).Methods(http.MethodGet)
}
