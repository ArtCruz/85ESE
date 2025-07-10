// Package classification Images API.
//
// Documentação da API de Imagens.
// Esta API permite upload, consulta e download de imagens associadas a produtos.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- multipart/form-data
//	- application/json
//
//	Produces:
//	- application/json
//	- image/png
//	- image/jpeg
//
// swagger:meta
package handlers

// swagger:response uploadResponse
type uploadResponseWrapper struct {
	// Mensagem de sucesso
	// in: body
	Message string `json:"message"`
}

// swagger:response errorResponse
type errorResponseWrapper struct {
	// Mensagem de erro
	// in: body
	Message string `json:"message"`
}

// swagger:parameters uploadImage
type uploadImageParamsWrapper struct {
	// ID do produto
	// in: formData
	// required: true
	ID string `json:"id"`
	// Arquivo de imagem
	// in: formData
	// required: true
	File string `json:"file"`
}

// swagger:parameters getImage
type getImageParamsWrapper struct {
	// ID do produto
	// in: path
	// required: true
	ID string `json:"id"`
}
