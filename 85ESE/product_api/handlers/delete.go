package handlers

import (
	"net/http"
	"product_api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletar um produto
//
// Remove um produto do sistema.
//
// parameters:
//   + name: id
//     in: path
//     description: ID do produto
//     required: true
//     type: integer
//
// responses:
//   204: noContentResponse
//   404: errorResponse
//   500: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	// err := data.DeleteProduct(id)
	err := p.repo.Delete(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
