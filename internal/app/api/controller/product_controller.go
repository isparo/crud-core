package controller

import (
	"crud-core/internal/app/api/controller/dto"
	"crud-core/internal/shared/errorhandler"
	"encoding/json"
	"log"
	"net/http"
)

type ProductService interface {
	Create(name string, price int) error
	List() ([]dto.Product, error)
}

type productHandler struct {
	prodService ProductService
}

func NewProductHandler(prodService ProductService) productHandler {
	return productHandler{
		prodService: prodService,
	}
}

func (ph productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("On create product request")

	var data dto.Product

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		errorMsg := errorhandler.NewValidationError("Bad request", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	log.Println(data)

	err = ph.prodService.Create(data.Name, data.Price)
	if err != nil {
		errorMsg := errorhandler.NewValidationError("Internal server error", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (ph productHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.prodService.List()
	if err != nil {
		errorMsg := errorhandler.NewValidationError("Internal server error", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(products)
}
