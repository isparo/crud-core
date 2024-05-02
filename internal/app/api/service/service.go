package service

import "crud-core/internal/app/api/model"

type ClientRepository interface {
	Create(client model.Client) error
	GetByID(id int) (*model.Client, error)
	Delete(id int) error
	Update(client model.Client) error
}

type clientService struct {
	clientRepository ClientRepository
}

func NewClientService(clientRepo ClientRepository) clientService {
	return clientService{
		clientRepository: clientRepo,
	}
}

func (cs clientService) Create(client model.Client) error {
	return nil
}
func (cs clientService) GetByID(id int) (*model.Client, error) {
	return nil, nil
}
func (cs clientService) Delete(id int) error {
	return nil
}
func (cs clientService) Update(client model.Client) error {
	return nil
}
