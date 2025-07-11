package data

import (
	"bytes"
	"testing"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.22,
		SKU:   "abc-efg-hji",
	}

	v := NewValidation()
	err := v.Validate(p)
	if len(err) == 0 {
		t.Error("Esperava erro de validação para produto sem nome")
	}
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name: "abc",
		SKU:  "abc-efg-hji",
	}

	v := NewValidation()
	err := v.Validate(p)
	if len(err) == 0 {
		t.Error("Esperava erro de validação para produto sem preço")
	}
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc",
	}

	v := NewValidation()
	err := v.Validate(p)
	if len(err) == 0 {
		t.Error("Esperava erro de validação para SKU inválido")
	}
}

func TestProductsToJSON(t *testing.T) { // testando dev direto main
	ps := []*Product{
		{
			Name:  "abc",
			Price: 1.22,
			SKU:   "abc-efg-hji",
		},
	}

	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	if err != nil {
		t.Errorf("Erro ao serializar produtos para JSON: %v", err)
	}
}
