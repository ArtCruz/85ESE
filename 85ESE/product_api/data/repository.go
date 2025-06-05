package data

type ProductRepository interface {
	GetAll() Products
	GetByID(id int) (*Product, error)
	Add(p Product) error
	Update(p Product) error
	Delete(id int) error
}
