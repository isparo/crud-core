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

type ProductHandler interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	ListProducts(w http.ResponseWriter, r *http.Request)
	GetProductByID(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type apiV1 struct {
	clientHandler clientHandler
	prodHandler   ProductHandler
}

func newAPIV1(clientHandler clientHandler, prodHandler ProductHandler) apiV1 {
	return apiV1{
		clientHandler: clientHandler,
		prodHandler:   prodHandler,
	}
}

func LoadAPI() {
	clientRepo := persistency.NewInMemoryRepository()

	clientService := service.NewClientService(clientRepo)
	clientHandler := controller.NewClientHandler(clientService)

	prodRepo := persistency.NewProductPersistency()
	prodService := service.NewProductService(prodRepo)
	prodHandler := controller.NewProductHandler(prodService)

	apiV1 := newAPIV1(clientHandler, prodHandler)
	apiV1.LoadRoutes()
}
