package data

type OrderRepository interface {
	GetAll() []*Order
	GetByID(id int) (*Order, error)
	Add(order Order) error
}
