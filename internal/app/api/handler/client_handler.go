package handler

import (
	"log"
	"net/http"
)

type clientHandler struct {
	//inject service here
}

func NewClientHandler() clientHandler {
	return clientHandler{}
}

func (ch clientHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("On create handler")
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
