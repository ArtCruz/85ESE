package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func UploadImage(apiURL string, r *http.Request) error {
	// Extrair o arquivo e o ID do produto da requisição
	file, header, err := r.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	id := r.FormValue("id")

	// Criar o corpo da requisição multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", header.Filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Adicionar o ID como campo do formulário
	err = writer.WriteField("id", id)
	if err != nil {
		return err
	}

	writer.Close()

	// Enviar a requisição para o serviço images
	req, err := http.NewRequest("POST", apiURL+"/upload", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Log da requisição enviada ao serviço images
	fmt.Printf("Enviando requisição para o serviço images: %s\n", apiURL)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Log da resposta recebida do serviço images
	fmt.Printf("Resposta do serviço images: %d %s\n", resp.StatusCode, resp.Status)

	return nil
}
