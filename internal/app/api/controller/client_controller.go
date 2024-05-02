package controller

import (
	"crud-core/internal/app/api/model"
	"log"
	"net/http"
)

type ClientService interface {
	Create(client model.Client) error
	GetByID(id int) (*model.Client, error)
	Delete(id int) error
	Update(client model.Client) error
}

type clientHandler struct {
	//inject service here
	clientService ClientService
}

func NewClientHandler(clientSvc ClientService) clientHandler {
	return clientHandler{
		clientService: clientSvc,
	}
}

func (ch clientHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("On create handler")

	//ch.clientService.Create()
}

func (ch clientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("On delete handler")

}

func (ch clientHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("On update handler")
}

func (ch clientHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("On get handler")
}

func (ch clientHandler) List(w http.ResponseWriter, r *http.Request) {
	log.Println("On List handler")
}
