package main

import (
	"testing"
)

// func TestOurClient(t *testing.T) {
// 	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
// 	c := client.NewHTTPClientWithConfig(nil, cfg)

// 	params := products.NewListProductsParams()
// 	prod, err := c.Products.ListProducts(params)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// fmt.Println(prod)

// 	fmt.Printf("%#v", prod.GetPayload()[0])
// 	t.Fail()
// }

func TestSomaSimples(t *testing.T) {
	if 1+1 != 2 {
		t.Error("Soma simples falhou")
	}
}
