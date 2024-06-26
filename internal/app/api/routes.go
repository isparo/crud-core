package api

import (
	"log"
	"net/http"
	"strings"
)

func (api apiV1) LoadRoutes() {
	log.Println("starting service")

	http.HandleFunc("/api/v1/clients", api.handlerClient)
	http.HandleFunc("/api/v1/clients/", api.handlerClient)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("can not start service")
	}
}

func (api apiV1) handlerClient(w http.ResponseWriter, r *http.Request) {

	pathValue := strings.TrimPrefix(r.URL.Path, "/api/v1/clients/")

	if !strings.Contains(pathValue, "/api/v1/clients") {
		switch r.Method {
		case http.MethodGet:
			api.clientHandler.Get(w, r)
		case http.MethodPut:
			api.clientHandler.Update(w, r)
		case http.MethodDelete:
			api.clientHandler.Delete(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	} else {
		switch r.Method {
		case http.MethodGet:
			api.clientHandler.List(w, r)
		case http.MethodPost:
			api.clientHandler.Create(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}
