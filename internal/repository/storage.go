package repository

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

type Storage interface {
	Read() (p []*Product, err error)
}

type StorageProduct interface {
	GetAll() (p []*Product, err error)
	GetByID(id int) (p *Product, err error)
	Search(price float64) (p []*Product, err error)
}
