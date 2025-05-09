package handlers

import (
	"gateway/config"
	"gateway/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := services.FetchProducts(cfg.ProductAPIURL)
		if err != nil {
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(products)
	}
}

func GetProduct(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := services.FetchProduct(cfg.ProductAPIURL, id)
		if err != nil {
			http.Error(w, "Erro ao buscar produto", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(product)
	}
}
