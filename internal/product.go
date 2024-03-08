package internal

import "errors"

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var (
	//ErrProductNotFound is returned when a product is not found in the database.
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	GetAll() []Product
	GetByID(id int) (Product, error)
	Search(price float64) []Product
}

type ProductService interface {
	GetAll() []Product
	GetByID(id int) (Product, error)
	Search(price float64) []Product
}
