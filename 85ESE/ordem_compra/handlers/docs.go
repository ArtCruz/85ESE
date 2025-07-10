// Package classification Ordem de Compra API.
//
// Documentação da API de Ordens de Compra.
// Esta API permite gerenciar ordens de compra, incluindo cadastro, consulta e listagem.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "ordem_compra/data"

// swagger:response ordersResponse
type ordersResponseWrapper struct {
	// Lista de ordens de compra
	// in: body
	Body []data.Order
}

// swagger:response orderResponse
type orderResponseWrapper struct {
	// Ordem de compra criada ou consultada
	// in: body
	Body data.Order
}

// swagger:parameters createOrder
type orderParamsWrapper struct {
	// Dados da ordem de compra para criação
	// in: body
	// required: true
	Body data.Order
}

// swagger:parameters getOrder
type orderIDParamsWrapper struct {
	// ID da ordem de compra
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response errorResponse
type errorResponseWrapper struct {
	// Mensagem de erro
	// in: body
	Body GenericError
}

// swagger:model
type GenericError struct {
	// Mensagem de erro
	Message string `json:"message"`
}
