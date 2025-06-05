package data

import (
	"fmt"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

type Products []*Product

func GetProducts() Products {
	GetLogger().Println("Listando todos os produtos")
	return productList
}

func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if i == -1 {
		GetLogger().Printf("Produto não encontrado para o ID: %d", id)
		return nil, ErrProductNotFound
	}
	GetLogger().Printf("Produto encontrado para o ID: %d", id)
	return productList[i], nil
}

func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		GetLogger().Printf("Erro ao atualizar: produto não encontrado para o ID: %d", p.ID)
		return ErrProductNotFound
	}
	productList[i] = &p
	GetLogger().Printf("Produto atualizado: %+v", p)
	return nil
}

func AddProduct(p Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
	GetLogger().Printf("Produto adicionado: %+v", p)
}

func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		GetLogger().Printf("Erro ao deletar: produto não encontrado para o ID: %d", id)
		return ErrProductNotFound
	}
	GetLogger().Printf("Produto deletado: %+v", *productList[i])
	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Cafezin Latte",
		Price:       2.45,
		SKU:         "abc323",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Cafezin Espresso",
		Price:       1.99,
		SKU:         "def456",
	},
}
