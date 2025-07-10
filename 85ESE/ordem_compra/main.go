package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"ordem_compra/data"
	"ordem_compra/handlers"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9092", "Bind address for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "orders-api ", log.LstdFlags)
	repo := data.NewOrderRepository("memory")
	oh := handlers.NewOrders(l, repo)

	sm := mux.NewRouter()

	// handlers para API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/orders", oh.ListAll)
	getR.HandleFunc("/orders/{id:[0-9]+}", oh.ListSingle)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/orders", oh.Create)

	// Servir documentação Swagger
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"http://localhost:8080"}),
		gohandlers.AllowedMethods([]string{"GET", "POST"}),
		gohandlers.AllowedHeaders([]string{"Content-Type"}),
	)

	s := http.Server{
		Addr:         *bindAddress,
		Handler:      ch(sm),
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9092")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

	orders := repo.GetAll()
	fmt.Println("Ordens:", orders)
}
