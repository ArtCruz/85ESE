package handlers

import (
	"bytes"
	"fmt"
	"gateway/config"
	"gateway/internal/services"
	"io"
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

	// Rota para o add products
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)

		// Lê o body para poder reutilizá-lo
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
			return
		}
		// Recria o body para uso posterior
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Faz a requisição manualmente para capturar status + body
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/products", cfg.ProductAPIURL), bytes.NewReader(bodyBytes))
		if err != nil {
			http.Error(w, "Erro ao criar requisição para o Product API", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Erro ao comunicar com Product API", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Erro ao ler resposta da Product API", http.StatusInternalServerError)
			return
		}

		// Repassa o status e corpo da resposta ao frontend
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(respBody)
	}).Methods(http.MethodPost)

	// Rota para atualizar um produto existente
	router.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)
		vars := mux.Vars(r)
		id := vars["id"]

		// Encaminha para o product_api
		updatedProduct, err := services.UpdateProduct(cfg.ProductAPIURL, id, r.Body)
		if err != nil {
			http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(updatedProduct)
	}).Methods(http.MethodPut)

	// Rota para deletar um produto existente
	router.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)
		vars := mux.Vars(r)
		id := vars["id"]

		// Encaminha para o product_api
		deletedProduct, err := services.DeleteProduct(cfg.ProductAPIURL, id)
		if err != nil {
			http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(deletedProduct)
	}).Methods(http.MethodDelete)

	// Rota para buscar uma imagem específica pelo ID
	router.HandleFunc("/images/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		// Monta a URL do serviço de imagens
		imageURL := fmt.Sprintf("%s/images/%s", cfg.ImagesAPIURL, id)

		// Faz o proxy da requisição GET para o serviço de imagens
		resp, err := http.Get(imageURL)
		if err != nil {
			http.Error(w, "Erro ao buscar imagem", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Copia o header de Content-Type
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}).Methods(http.MethodGet)
}
