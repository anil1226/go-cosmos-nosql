package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/anil1226/go-employee/internal/models"
	"github.com/gorilla/mux"
)

type EmpService interface {
	GetEmployee(context.Context, string) (models.Employee, error)
	CreateEmployee(context.Context, models.Employee) error
	UpdateEmployee(context.Context, models.Employee) error
	DeleteEmployee(context.Context, string) error
}

type Response struct {
	Message string
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var postCmt models.Employee
	if err := json.NewDecoder(r.Body).Decode(&postCmt); err != nil {
		return
	}

	err := h.Service.CreateEmployee(r.Context(), postCmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Created"}); err != nil {
		panic(err)
	}

}

// @Summary Returns a employee json
// @Description Returns a employee json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /api/v1/employee/{id} [get]
func (h *Handler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetEmployee(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}

func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	var cmt models.Employee
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	err := h.Service.UpdateEmployee(r.Context(), cmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Updated"}); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteEmployee(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}
