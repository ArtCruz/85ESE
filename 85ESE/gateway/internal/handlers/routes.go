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
	router.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)

		uploader := &services.HTTPImageUploader{APIURL: cfg.ImagesAPIURL}
		err := uploader.UploadImage(r)
		if err != nil {
			fmt.Printf("Erro ao enviar imagem para o serviço images: %v\n", err)
			http.Error(w, "Erro ao enviar imagem", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Imagem enviada com sucesso"))
	}).Methods(http.MethodPost)

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)

		products, err := services.FetchProducts(cfg.ProductAPIURL)
		if err != nil {
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(products)
	}).Methods(http.MethodGet)

	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(respBody)
	}).Methods(http.MethodPost)

	router.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)
		vars := mux.Vars(r)
		id := vars["id"]

		updatedProduct, err := services.UpdateProduct(cfg.ProductAPIURL, id, r.Body)
		if err != nil {
			http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(updatedProduct)
	}).Methods(http.MethodPut)

	router.HandleFunc("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requisição recebida do frontend: %s %s\n", r.Method, r.URL.Path)
		vars := mux.Vars(r)
		id := vars["id"]

		deletedProduct, err := services.DeleteProduct(cfg.ProductAPIURL, id)
		if err != nil {
			http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(deletedProduct)
	}).Methods(http.MethodDelete)

	router.HandleFunc("/images/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		imageURL := fmt.Sprintf("%s/images/%s", cfg.ImagesAPIURL, id)

		resp, err := http.Get(imageURL)
		if err != nil {
			http.Error(w, "Erro ao buscar imagem", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}).Methods(http.MethodGet)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/orders", cfg.OrdersAPIURL)
		var resp *http.Response
		var err error

		if r.Method == http.MethodGet {
			resp, err = http.Get(url)
		} else if r.Method == http.MethodPost {
			bodyBytes, _ := io.ReadAll(r.Body)
			resp, err = http.Post(url, "application/json", bytes.NewReader(bodyBytes))
		}

		if err != nil {
			http.Error(w, "Erro ao comunicar com ordem_compra", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}).Methods(http.MethodGet, http.MethodPost)

	router.HandleFunc("/login.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /login.html proxy para frontend")

		resp, err := http.Get(fmt.Sprintf("%s/login.html", cfg.FrontendURL))
		if err != nil {
			http.Error(w, "Erro ao buscar página de login", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}).Methods(http.MethodGet)

	router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Veio pro GATEWAY")

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf("%s/auth", cfg.AuthAPIURL)
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
		if err != nil {
			http.Error(w, "Erro ao criar requisição para Auth API", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Erro ao comunicar com Auth API", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}).Methods("POST")
}
