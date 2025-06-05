package services

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Busca todos os produtos do product_api
func FetchProducts(apiURL string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products", apiURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Busca um produto específico pelo ID
func FetchProduct(apiURL, id string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", apiURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Cria um novo produto
func CreateProduct(apiURL string, body io.Reader) ([]byte, error) {
	// Lê o conteúdo do body para debug
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Conteúdo enviado para o product_api:", string(data))

	// Cria um novo reader para enviar no POST
	resp, err := http.Post(fmt.Sprintf("%s/products", apiURL), "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Atualiza um produto existente
func UpdateProduct(apiURL string, id string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/products/%s", apiURL, id), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Deleta um produto
func DeleteProduct(apiURL string, id string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/products/%s", apiURL, id), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
