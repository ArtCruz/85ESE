package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ordem_compra/data"
)

func CreateOrder(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Recebida requisição para criar nova ordem de compra")
	var order data.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		http.Error(rw, "Dados inválidos", http.StatusBadRequest)
		return
	}
	if order.ProductID == 0 || order.Quantity <= 0 {
		fmt.Println("ProdutoID ou quantidade inválidos:", order)
		http.Error(rw, "Produto e quantidade obrigatórios", http.StatusBadRequest)
		return
	}
	fmt.Printf("Criando ordem: %+v\n", order)
	data.AddOrder(order)
	fmt.Println("Ordem criada com sucesso!")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(order)
}
