package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service EmpService
	Router  *mux.Router
	Server  *http.Server
}

func NewHandler(service EmpService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {

	h.Router.HandleFunc("/api/v1/employee", h.CreateEmployee).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/v1/employee/{id}", h.GetEmployee).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/v1/employee", h.UpdateEmployee).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/v1/employee/{id}", h.DeleteEmployee).Methods(http.MethodDelete)
}

func (h *Handler) Serve() error {

	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Server.Shutdown(context)
	return nil
}
