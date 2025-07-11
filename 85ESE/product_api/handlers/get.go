package handlers

import (
	"fmt"
	"net/http"
	"product_api/data"
)

// swagger:route GET /products products listProducts
// Listar todos os produtos
//
// Retorna uma lista de todos os produtos cadastrados.
//
// responses:
//   200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	rw.Header().Add("Content-Type", "application/json")

	fmt.Printf("Requisição recebida do gateway: %s %s\n", r.Method, r.URL.Path)

	prods := p.repo.GetAll()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}

}

// swagger:route GET /products/{id} products listSingleProduct
// Buscar produto por ID
//
// Retorna os detalhes de um produto específico.
//
// parameters:
//   + name: id
//     in: path
//     description: ID do produto
//     required: true
//     type: integer
//
// responses:
//   200: productResponse
//   404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
