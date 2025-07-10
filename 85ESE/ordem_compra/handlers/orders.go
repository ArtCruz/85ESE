package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"ordem_compra/data"
	"strconv"

	"github.com/gorilla/mux"
)

// Orders
// swagger:route GET /orders orders listOrders
// Listar todas as ordens de compra
//
// Retorna uma lista de todas as ordens de compra.
//
// responses:
//
//	200: ordersResponse
type Orders struct {
	l    *log.Logger
	repo data.OrderRepository
}

func NewOrders(l *log.Logger, repo data.OrderRepository) *Orders {
	return &Orders{l, repo}
}

func (o *Orders) ListAll(rw http.ResponseWriter, r *http.Request) {
	orders := o.repo.GetAll()
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(orders)
}

// swagger:route GET /orders/{id} orders getOrder
// Buscar ordem de compra por ID
//
// Retorna os detalhes de uma ordem de compra específica.
//
// parameters:
//   - name: id
//     in: path
//     description: ID da ordem de compra
//     required: true
//     type: integer
//
// responses:
//
//	200: orderResponse
//	404: errorResponse
func (o *Orders) ListSingle(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "ID inválido", http.StatusBadRequest)
		return
	}
	order, err := o.repo.GetByID(id)
	if err != nil {
		http.Error(rw, "Ordem não encontrada", http.StatusNotFound)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(order)
}

// swagger:route POST /orders orders createOrder
// Criar uma nova ordem de compra
//
// Adiciona uma nova ordem de compra ao sistema.
//
// responses:
//
//	201: orderResponse
//	400: errorResponse
func (o *Orders) Create(rw http.ResponseWriter, r *http.Request) {
	var order data.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(rw, "Dados inválidos", http.StatusBadRequest)
		return
	}
	if order.ProductID == 0 || order.Quantity <= 0 {
		http.Error(rw, "Produto e quantidade obrigatórios", http.StatusBadRequest)
		return
	}
	o.repo.Add(order)
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(order)
}
