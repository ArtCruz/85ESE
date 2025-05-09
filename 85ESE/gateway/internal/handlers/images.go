package handlers

import (
	"gateway/config"
	"gateway/internal/services"
	"net/http"
)

func UploadImage(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := services.UploadImage(cfg.ImagesAPIURL, r)
		if err != nil {
			http.Error(w, "Erro ao fazer upload da imagem", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Imagem enviada com sucesso"))
	}
}
