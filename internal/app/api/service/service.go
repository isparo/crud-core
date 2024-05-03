package service

import (
	"crud-core/internal/app/api/model"
	"log"
)

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

func (cs clientService) Create(name string, email string) error {
	log.Println("on create service")

	client := model.NewClient(nil, name, email)
	cs.clientRepository.Create(client)

	return nil
}
func (cs clientService) GetByID(id int) (*model.Client, error) {

	return cs.clientRepository.GetByID(id)
}
func (cs clientService) Delete(id int) error {
	return nil
}
func (cs clientService) Update(id int, name string, email string) error {

	client := model.NewClient(&id, name, email)

	return cs.clientRepository.Update(client)
}
