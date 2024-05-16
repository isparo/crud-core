package model

type Product struct {
	ID    *int
	Name  string
	Price int
}

func NewProduct(name string, price int) Product {
	return Product{
		ID:    nil,
		Name:  name,
		Price: price,
	}
}
