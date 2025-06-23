package handlers

import (
	"gateway/config"
	"gateway/internal/services"
	"net/http"
)

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
//     proxy := httputil.NewSingleHostReverseProxy(&url.URL{
//         Scheme: "http",
//         Host:   "localhost:7070",
//     })
//     proxy.ServeHTTP(w, r)
// }

func Auth(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Passa o corpo da requisição original (r.Body) para o serviço de login
		respBody, err := services.Login(cfg.AuthAPIURL, r.Body)
		if err != nil {
			http.Error(w, "Erro ao fazer autenticação", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(respBody)
	}
}
