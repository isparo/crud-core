package controller

import (
	"crud-core/internal/app/api/controller/dto"
	"crud-core/internal/app/api/model"
	"encoding/json"
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
	var data dto.Client

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		//http.Error(w, "Error al decodificar el cuerpo de la solicitud JSON", http.StatusBadRequest)
		errorMsg := struct {
			Message string `json:"message"`
			Code    int    `json:"status"`
		}{
			Message: "Error al decodificar el cuerpo de la solicitud JSON",
			Code:    http.StatusBadRequest,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	log.Println(data)

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
