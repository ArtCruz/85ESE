package handlers

import (
    "net/http"
    "net/http/httputil"
    "net/url"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    proxy := httputil.NewSingleHostReverseProxy(&url.URL{
        Scheme: "http",
        Host:   "localhost:7070",
    })
    proxy.ServeHTTP(w, r)
}
