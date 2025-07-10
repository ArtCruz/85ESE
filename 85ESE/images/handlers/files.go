package handlers

import (
	"fmt"
	"images/files"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// UploadREST implements the http.Handler interface
// swagger:route POST /upload images uploadImage
//
// # Fazer upload de uma imagem para um produto
//
// Envia uma imagem associada a um produto.
//
// Consome multipart/form-data.
//
// responses:
//
//	200: uploadResponse
//	400: errorResponse
//	500: errorResponse
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters
	f.saveFile(id, fn, rw, r.Body)
}

func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	// Log da requisição recebida do gateway
	fmt.Printf("Requisição recebida do gateway: %s %s\n", r.Method, r.URL.Path)

	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Erro ao processar formulário multipart", "error", err)
		http.Error(rw, "Erro ao processar formulário multipart", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	file, header, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Erro ao obter arquivo", "error", err)
		http.Error(rw, "Erro ao obter arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Salvar o arquivo
	f.saveFile(id, header.Filename, rw, file)
}

func (f *Files) Ping(rw http.ResponseWriter, r *http.Request) {
	// Log da requisição de ping
	fmt.Println("Ping recebido no serviço images")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Pong"))
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Salvando arquivo para o produto", "id", id, "path", path)

	// Caminho da pasta do produto
	dir := filepath.Join("imagestore", id)

	// Remove todas as imagens antigas antes de salvar a nova
	files, err := os.ReadDir(dir)
	if err == nil {
		for _, file := range files {
			if !file.IsDir() && (strings.HasSuffix(file.Name(), ".png") || strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg")) {
				os.Remove(filepath.Join(dir, file.Name()))
			}
		}
	}

	fp := filepath.Join(id, path)
	err = f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Erro ao salvar arquivo", "error", err)
		http.Error(rw, "Erro ao salvar arquivo", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Arquivo salvo com sucesso"))
}

type LocalStorage struct {
	Path   string
	Logger hclog.Logger
}

// Save implements files.Storage.
func (l *LocalStorage) Save(path string, file io.Reader) error {
	// Caminho completo para o arquivo
	fullPath := filepath.Join(l.Path, path)

	// Criar diretório, se necessário
	dir := filepath.Dir(fullPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		l.Logger.Info("Criando diretório", "path", dir)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			l.Logger.Error("Erro ao criar diretório", "path", dir, "error", err)
			return err
		}
	}

	// Criar o arquivo
	l.Logger.Info("Salvando arquivo", "path", fullPath)
	out, err := os.Create(fullPath)
	if err != nil {
		l.Logger.Error("Erro ao criar arquivo", "path", fullPath, "error", err)
		return err
	}
	defer out.Close()

	// Copiar o conteúdo do arquivo recebido para o arquivo criado
	_, err = io.Copy(out, file)
	if err != nil {
		l.Logger.Error("Erro ao salvar conteúdo no arquivo", "path", fullPath, "error", err)
		return err
	}

	l.Logger.Info("Arquivo salvo com sucesso", "path", fullPath)
	return nil
}

func NewLocalStorage(path string, logger hclog.Logger) *LocalStorage {
	// Ensure the directory exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			logger.Error("Failed to create directory", "path", path, "error", err)
			return nil
		}
	}

	return &LocalStorage{
		Path:   path,
		Logger: logger,
	}
}

func (f *Files) ServeProductImage(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dir := filepath.Join("imagestore", id)
	files, err := os.ReadDir(dir)
	if err != nil || len(files) == 0 {
		http.NotFound(rw, r)
		return
	}

	// Pega o primeiro arquivo de imagem
	fmt.Println("Busca Imagem com ID:", id)
	for _, file := range files {
		if !file.IsDir() && (strings.HasSuffix(file.Name(), ".png") || strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg")) {
			imgPath := filepath.Join(dir, file.Name())
			http.ServeFile(rw, r, imgPath)
			return
		}
	}

	http.NotFound(rw, r)
}

// swagger:route GET /images/{id} images getImage
// Obter imagem de um produto
//
// Retorna a imagem associada ao produto.
//
// parameters:
//   + name: id
//     in: path
//     description: ID do produto
//     required: true
//     type: string
//
// produces:
//   - image/png
//   - image/jpeg
//
// responses:
//   200: file
//   404: errorResponse

// swagger:route GET /ping images ping
// Healthcheck do serviço de imagens
//
// Verifica se o serviço está online.
//
// responses:
//   200: uploadResponse
