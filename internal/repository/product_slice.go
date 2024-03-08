package repository

func NewProductSlice(st Storage) *ProductSlice {
	return &ProductSlice{st: st}
}

type ProductSlice struct {
	st Storage
}

func (p *ProductSlice) GetAll() (products []*Product, err error) {
	products, err = p.st.Read()
	if err != nil {
		return nil, err
	}

	return
}
