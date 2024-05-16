package dto

type Product struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"proce"`
}

func NewProduct(
	id *int,
	name string,
	price int,
) Product {
	return Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
