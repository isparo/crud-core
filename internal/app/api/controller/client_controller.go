package controller

import (
	"crud-core/internal/app/api/controller/dto"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ClientService interface {
	Create(name string, email string) error
	GetByID(id int) (*dto.Client, error)
	Delete(id int) error
	Update(id int, name string, email string) error
	GetClients() ([]dto.Client, error)
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

	ch.clientService.Create(data.Name, data.Email)
	w.WriteHeader(http.StatusAccepted)
}

func (ch clientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("On delete handler")

	idPath := strings.TrimPrefix(r.URL.Path, "/api/v1/clients/")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		errorMsg := struct {
			Message string `json:"message"`
			Code    int    `json:"status"`
		}{
			Message: "wrong id value",
			Code:    http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	err = ch.clientService.Delete(id)
	if err != nil {
		errorMsg := struct {
			Message string `json:"message"`
			Code    int    `json:"status"`
		}{
			Message: "Internal eror ",
			Code:    http.StatusInternalServerError,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}

func (ch clientHandler) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("On update handler")

	idPath := strings.TrimPrefix(r.URL.Path, "/api/v1/clients/")
	id, err := strconv.Atoi(idPath)

	if err != nil {
		errorMsg := struct {
			Message string `json:"message"`
			Code    int    `json:"status"`
		}{
			Message: "wrong id value",
			Code:    http.StatusBadRequest,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	var data dto.Client
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ch.clientService.Update(id, data.Name, data.Email)

	w.WriteHeader(http.StatusAccepted)
}

func (ch clientHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("On get handler")

	idPath := strings.TrimPrefix(r.URL.Path, "/api/v1/clients/")

	id, err := strconv.Atoi(idPath)
	if err != nil {
		errorMsg := struct {
			Message string `json:"message"`
			Code    int    `json:"status"`
		}{
			Message: "wrong id value",
			Code:    http.StatusBadRequest,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	log.Println("searching: ", id)

	client, err := ch.clientService.GetByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(client)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

func (ch clientHandler) List(w http.ResponseWriter, r *http.Request) {
	log.Println("On List handler")

	clients, err := ch.clientService.GetClients()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(clients)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}
