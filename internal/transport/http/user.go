package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/anil1226/go-employee/internal/models"
)

type UserService interface {
	GetUser(context.Context, string) (models.User, error)
	CreateUser(context.Context, models.User) error
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var postCmt models.User
	if err := json.NewDecoder(r.Body).Decode(&postCmt); err != nil {
		return
	}

	err := h.Service.CreateUser(r.Context(), postCmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Created"}); err != nil {
		panic(err)
	}

}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	var postCmt models.User
	if err := json.NewDecoder(r.Body).Decode(&postCmt); err != nil {
		return
	}

	if postCmt.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetUser(r.Context(), postCmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := GenerateJWT(cmt.Name)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := models.UserResp{
		Name:  cmt.Name,
		ID:    cmt.ID,
		Token: token,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}

}
