package data

func NewOrderRepository(repoType string) OrderRepository {
	switch repoType {
	case "memory":
		return NewMemoryOrderRepository()
	default:
		return NewMemoryOrderRepository()
	}
}
