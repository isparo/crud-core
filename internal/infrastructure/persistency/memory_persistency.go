package persistency

import (
	"crud-core/internal/domain/client"
	"errors"
	"sync"
	"time"
)

type inMemoryRepository struct {
	dataStorage map[int]client.Client
	mtx         sync.Mutex
}

func NewInMemoryRepository() inMemoryRepository {
	return inMemoryRepository{
		dataStorage: make(map[int]client.Client),
	}
}

func (r *inMemoryRepository) Create(client client.Client) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	client.ID = int(time.Now().UnixNano())
	r.dataStorage[client.ID] = client

	return nil
}

func (r *inMemoryRepository) GetByID(id int) (*client.Client, error) {
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

func (r *inMemoryRepository) Update(client client.Client) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, ok := r.dataStorage[client.ID]; !ok {
		return errors.New("client not found")
	}

	r.dataStorage[client.ID] = client

	return nil
}
