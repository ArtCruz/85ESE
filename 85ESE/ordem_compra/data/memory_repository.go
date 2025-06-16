package data

import "fmt"

type MemoryOrderRepository struct {
	orders []*Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: []*Order{},
	}
}

func (m *MemoryOrderRepository) GetAll() []*Order {
	return m.orders
}

func (m *MemoryOrderRepository) GetByID(id int) (*Order, error) {
	for _, o := range m.orders {
		if o.ID == id {
			return o, nil
		}
	}
	return nil, fmt.Errorf("Order not found")
}

func (m *MemoryOrderRepository) Add(order Order) error {
	order.ID = len(m.orders) + 1
	m.orders = append(m.orders, &order)
	fmt.Printf("Order adicionada ao repositório em memória: %+v\n", order)
	return nil
}
