package data

func NewProductRepository(repoType string) ProductRepository {
	switch repoType {
	case "memory":
		return NewMemoryProductRepository()
	// No futuro, adicione outros tipos, como "postgres", "mongo", etc.
	default:
		return NewMemoryProductRepository()
	}
}
