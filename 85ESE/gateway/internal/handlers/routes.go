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

	// Rota para autenticação: login
	// router.HandleFunc("/login.html", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("GET /login.html proxy para frontend")

	// 	resp, err := http.Get(fmt.Sprintf("%s/login.html", cfg.FrontendURL))
	// 	if err != nil {
	// 		http.Error(w, "Erro ao buscar página de login", http.StatusBadGateway)
	// 		return
	// 	}
	// 	defer resp.Body.Close()

	// 	for k, vv := range resp.Header {
	// 		for _, v := range vv {
	// 			w.Header().Add(k, v)
	// 		}
	// 	}
	// 	w.WriteHeader(resp.StatusCode)
	// 	io.Copy(w, resp.Body)
	// }).Methods(http.MethodGet)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		// Proxy GET e POST para o serviço ordem_compra
		url := fmt.Sprintf("%s/orders", cfg.OrdersAPIURL) // Adicione OrdersAPIURL no config.go
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

	// Rota para autenticação: login
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

	router.HandleFunc("/gateway/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Proxy: POST /gateway/auth para serviço de autenticação")

		// Recria o body para reutilização
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Cria a requisição para o AuthAPI
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth", cfg.AuthAPIURL), bytes.NewBuffer(bodyBytes))
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

		// Copia a resposta para o cliente (frontend)
		copyResponse(w, resp)
	}).Methods(http.MethodPost)

	router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("AAAAAAAAAAAAAIIIIIIIIIIIIII")
		url := fmt.Sprintf("%s/auth", cfg.AuthAPIURL)
		req, err := http.NewRequest(http.MethodPost, url, r.Body)
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

func forwardRequest(originalReq *http.Request, url string) (*http.Response, error) {
	body, err := io.ReadAll(originalReq.Body)
	if err != nil {
		return nil, err
	}

	newReq, err := http.NewRequest(originalReq.Method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Copia os headers
	for name, values := range originalReq.Header {
		for _, value := range values {
			newReq.Header.Add(name, value)
		}
	}

	client := &http.Client{}
	return client.Do(newReq)
}

// Copia a resposta do serviço para o ResponseWriter
func copyResponse(w http.ResponseWriter, resp *http.Response) error {
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, err := io.Copy(w, resp.Body)
	return err
}
