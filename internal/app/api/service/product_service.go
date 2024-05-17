package service

import (
	"crud-core/internal/app/api/controller/dto"
	"crud-core/internal/app/api/model"
	"log"
)

type ProductRepository interface {
	Add(product model.Product) error
	List() []model.Product
	GetByID(id int) (*model.Product, error)
	Delete(id int) error
}

type productService struct {
	prodRepo ProductRepository
}

func NewProductService(prodRepo ProductRepository) productService {
	return productService{
		prodRepo: prodRepo,
	}
}

func (ps productService) Create(name string, price int) error {
	log.Println("On create product service")

	err := ps.prodRepo.Add(model.NewProduct(name, price))
	if err != nil {
		return err
	}

	return nil
}

func (ps productService) List() ([]dto.Product, error) {

	products := []dto.Product{}

	prods := ps.prodRepo.List()

	for _, p := range prods {
		products = append(products, dto.NewProduct(p.ID, p.Name, p.Price))
	}

	return products, nil
}

func (ps productService) GetByID(id int) (*dto.Product, error) {

	prod, err := ps.prodRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	product := dto.NewProduct(prod.ID, prod.Name, prod.Price)

	return &product, nil
}

func (ps productService) DeleteClient(id int) error {
	err := ps.prodRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
