package persistency

import (
	"crud-core/internal/app/api/model"
	"errors"
	"log"
	"sync"
	"time"
)

type inMemoryRepository struct {
	dataStorage map[int]model.Client
	mtx         sync.Mutex
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		dataStorage: make(map[int]model.Client),
	}
}

func (r *inMemoryRepository) Create(client model.Client) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	id := int(time.Now().UnixNano())
	client.ID = &id

	log.Println("created: ", *client.ID)

	r.dataStorage[*client.ID] = client

	return nil
}

func (r *inMemoryRepository) GetByID(id int) (*model.Client, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	client, ok := r.dataStorage[id]

	if !ok {
		return nil, errors.New("client not found")
	}

	return &client, nil
}

func (r *inMemoryRepository) Delete(id int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.dataStorage[id]; !ok {
		return errors.New("client not found")
	}

	delete(r.dataStorage, id)

	return nil
}

func (r *inMemoryRepository) Update(client model.Client) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.dataStorage[*client.ID]; !ok {
		return errors.New("client not found")
	}

	r.dataStorage[*client.ID] = client

	return nil
}
