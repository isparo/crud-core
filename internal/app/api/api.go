package api

import (
	"crud-core/internal/app/api/controller"
	"crud-core/internal/app/api/service"
	"crud-core/internal/infrastructure/persistency"
	"net/http"
)

type clientHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type apiV1 struct {
	clientHandler clientHandler
}

func newAPIV1(clientHandler clientHandler) apiV1 {
	return apiV1{
		clientHandler: clientHandler,
	}
}

func LoadAPI() {
	clientRepo := persistency.NewInMemoryRepository()

	clientService := service.NewClientService(clientRepo)
	clientHandler := controller.NewClientHandler(clientService)
	apiV1 := newAPIV1(clientHandler)
	apiV1.LoadRoutes()
}
