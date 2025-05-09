package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchProducts(apiURL string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products", apiURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func FetchProduct(apiURL, id string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", apiURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
