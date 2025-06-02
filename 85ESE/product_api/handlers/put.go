package handlers

import (
	"net/http"
	"product_api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "ID inválido", http.StatusBadRequest)
		return
	}

	// Pegue o produto do contexto, preenchido pelo middleware
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	prod.ID = id

	err = data.UpdateProduct(*prod)
	if err != nil {
		http.Error(rw, "Produto não encontrado", http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	data.ToJSON(prod, rw)
}
