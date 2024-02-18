package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/anil1226/go-employee/docs"
	"github.com/anil1226/go-employee/internal/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	Service *service.Service
	Router  *mux.Router
	Server  *http.Server
}

func NewHandler(service *service.Service) *Handler {
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func (h *Handler) mapRoutes() {

	h.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	h.Router.HandleFunc("/api/v1/employee", verifyJWT(h.CreateEmployee)).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/v1/employee/{id}", verifyJWT(h.GetEmployee)).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/v1/employee", verifyJWT(h.UpdateEmployee)).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/v1/employee/{id}", verifyJWT(h.DeleteEmployee)).Methods(http.MethodDelete)

	h.Router.HandleFunc("/api/v1/user", h.CreateUser).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/v1/signin", h.GetUser).Methods(http.MethodPost)
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
