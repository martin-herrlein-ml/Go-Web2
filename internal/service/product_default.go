package service

import "GoWeb/proyect/internal"

func newDefaultProduct(repository internal.ProductRepository) *ProductDefault {
	return &ProductDefault{repository: repository}
}

type ProductDefault struct {
	repository internal.ProductRepository
}

func (p *ProductDefault) GetAll() ([]internal.Product, error) {
	return p.repository.GetAll(), nil
}
