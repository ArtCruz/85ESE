package handlers

import (
	// "gateway/config"
	// "gateway/internal/services"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func LoginPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remote, err := url.Parse("http://localhost:7071")
		if err != nil {
			http.Error(w, "Erro no proxy da tela de login", http.StatusInternalServerError)
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		r.URL.Path = "/login.html"
		r.Host = remote.Host
		proxy.ServeHTTP(w, r)
	}
}
