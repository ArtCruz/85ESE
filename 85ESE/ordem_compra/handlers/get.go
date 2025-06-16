package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ordem_compra/data"
)

func ListOrders(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebida requisição para listar ordens de compra")
	orders := data.GetOrders()
	fmt.Printf("Retornando %d ordens\n", len(orders))
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(orders)
}
