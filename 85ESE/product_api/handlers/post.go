package handlers

import (
	"net/http"
	"product_api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	err := p.repo.Add(*prod)
	if err != nil {
		http.Error(rw, "Erro ao adicionar produto", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	data.ToJSON(prod, rw)
}
