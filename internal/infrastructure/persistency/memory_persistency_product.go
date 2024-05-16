package persistency

import (
	"crud-core/internal/app/api/model"
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
