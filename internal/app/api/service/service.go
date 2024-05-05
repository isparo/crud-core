package service

import (
	"crud-core/internal/app/api/controller/dto"
	"crud-core/internal/app/api/model"
	"log"
)

type ClientRepository interface {
	Create(client model.Client) error
	GetByID(id int) (*model.Client, error)
	Delete(id int) error
	Update(client model.Client) error
	List() ([]model.Client, error)
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
func (cs clientService) GetByID(id int) (*dto.Client, error) {

	client, err := cs.clientRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.Client{client.ID, client.Name, client.Email}, nil
}
func (cs clientService) Delete(id int) error {

	if err := cs.clientRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
func (cs clientService) Update(id int, name string, email string) error {

	client := model.NewClient(&id, name, email)

	return cs.clientRepository.Update(client)
}

func (cs clientService) GetClients() ([]dto.Client, error) {

	clientsResp := []dto.Client{}
	clients, err := cs.clientRepository.List()
	if err != nil {
		return clientsResp, err
	}

	for _, v := range clients {
		clientsResp = append(clientsResp, dto.NewClient(v.ID, v.Name, v.Email))
	}

	return clientsResp, nil
}
