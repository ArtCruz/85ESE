package data

import "fmt"

type Order struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
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
