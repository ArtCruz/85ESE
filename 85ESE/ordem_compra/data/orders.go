package data

import "fmt"

// Ordem de Compra representa um pedido realizado por um cliente.
// swagger:model
type Order struct {
	// ID da ordem de compra
	ID int `json:"id"`
	// ID do produto
	ProductID int `json:"product_id"`
	// Quantidade do produto
	Quantity int `json:"quantity"`
}

var orders []*Order

var ErrOrderNotFound = fmt.Errorf("Order not found")

func AddOrder(order Order) {
	order.ID = len(orders) + 1
	orders = append(orders, &order)
}

func GetOrders() []*Order {
	return orders
}

func GetOrderByID(id int) (*Order, error) {
	for _, o := range orders {
		if o.ID == id {
			return o, nil
		}
	}
	return nil, ErrOrderNotFound
}
