package services

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Login(apiURL string, body io.Reader) ([]byte, error) {
	// Lê o conteúdo do body para debug
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Conteúdo enviado para a autenticação:", string(data))

	// Cria um novo reader para enviar no POST
	resp, err := http.Post(fmt.Sprintf("%s/auth", apiURL), "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
