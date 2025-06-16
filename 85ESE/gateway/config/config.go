package config

import (
	"github.com/nicholasjackson/env"
)

type Config struct {
	ServerAddress string
	ProductAPIURL string
	ImagesAPIURL  string
	OrdersAPIURL  string
}

func Load() *Config {
	serverAddress := env.String("GATEWAY_ADDRESS", false, ":8080", "Endereço do gateway")
	productAPIURL := env.String("PRODUCT_API_URL", false, "http://localhost:9090", "URL do serviço product_api")
	imagesAPIURL := env.String("IMAGES_API_URL", false, "http://localhost:9091", "URL do serviço images")
	ordersAPIURL := env.String("ORDERS_API_URL", false, "http://localhost:9092", "URL do serviço ordem_compra")

	env.Parse()

	return &Config{
		ServerAddress: *serverAddress,
		ProductAPIURL: *productAPIURL,
		ImagesAPIURL:  *imagesAPIURL,
		OrdersAPIURL:  *ordersAPIURL,
	}
}
