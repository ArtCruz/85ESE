package data

type MemoryProductRepository struct {
	products Products
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: productList,
	}
}

func (m *MemoryProductRepository) GetAll() Products {
	GetLogger().Println("Listando todos os produtos (Memory)")
	return m.products
}

func (m *MemoryProductRepository) GetByID(id int) (*Product, error) {
	for _, p := range m.products {
		if p.ID == id {
			GetLogger().Printf("Produto encontrado para o ID: %d (Memory)", id)
			return p, nil
		}
	}
	GetLogger().Printf("Produto não encontrado para o ID: %d (Memory)", id)
	return nil, ErrProductNotFound
}

func (m *MemoryProductRepository) Add(p Product) error {
	maxID := 0
	if len(m.products) > 0 {
		maxID = m.products[len(m.products)-1].ID
	}
	p.ID = maxID + 1
	m.products = append(m.products, &p)
	GetLogger().Printf("Produto adicionado: %+v (Memory)", p)
	return nil
}

func (m *MemoryProductRepository) Update(p Product) error {
	for i, prod := range m.products {
		if prod.ID == p.ID {
			m.products[i] = &p
			GetLogger().Printf("Produto atualizado: %+v (Memory)", p)
			return nil
		}
	}
	GetLogger().Printf("Erro ao atualizar: produto não encontrado para o ID: %d (Memory)", p.ID)
	return ErrProductNotFound
}

func (m *MemoryProductRepository) Delete(id int) error {
	for i, prod := range m.products {
		if prod.ID == id {
			GetLogger().Printf("Produto deletado: %+v (Memory)", *prod)
			m.products = append(m.products[:i], m.products[i+1:]...)
			return nil
		}
	}
	GetLogger().Printf("Erro ao deletar: produto não encontrado para o ID: %d (Memory)", id)
	return ErrProductNotFound
}
