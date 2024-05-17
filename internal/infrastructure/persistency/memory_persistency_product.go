package persistency

import (
	"crud-core/internal/app/api/model"
	"errors"
	"log"
	"sync"
	"time"
)

type productPersistency struct {
	dataStorage map[int]model.Product
	mtx         sync.Mutex
}

func NewProductPersistency() *productPersistency {
	return &productPersistency{
		dataStorage: make(map[int]model.Product),
	}
}

func (pp *productPersistency) Add(product model.Product) error {
	pp.mtx.Lock()
	defer pp.mtx.Unlock()

	id := int(time.Now().UnixNano())
	product.ID = &id
	pp.dataStorage[id] = product

	log.Println("product added with id: ", *product.ID)

	return nil
}

func (pp *productPersistency) List() []model.Product {
	pp.mtx.Lock()
	defer pp.mtx.Unlock()

	products := []model.Product{}

	for _, v := range pp.dataStorage {
		products = append(products, v)
	}

	return products
}

func (pp *productPersistency) GetByID(id int) (*model.Product, error) {
	pp.mtx.Lock()
	defer pp.mtx.Unlock()

	if _, ok := pp.dataStorage[id]; !ok {
		return nil, errors.New("product not found")
	}

	prod := pp.dataStorage[id]

	return &prod, nil
}
func (pp *productPersistency) Delete(id int) error {
	pp.mtx.Lock()
	defer pp.mtx.Unlock()

	if _, ok := pp.dataStorage[id]; !ok {
		return errors.New("client not found")
	}

	delete(pp.dataStorage, id)

	return nil
}

func (pp *productPersistency) Update(product model.Product) error {
	pp.mtx.Lock()
	defer pp.mtx.Unlock()

	if _, ok := pp.dataStorage[*product.ID]; !ok {
		return errors.New("product not found")
	}

	pp.dataStorage[*product.ID] = product

	return nil
}
