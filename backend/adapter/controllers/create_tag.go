package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/h4shu/lounge-books/adapter/presenters"
	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/usecases"
)

type (
	CreateTagController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}
	createTagController struct {
		usecase   usecases.CreateTagUsecase
		presenter presenters.CreateTagPresenter
	}
	createTagRequest struct {
		Name string `json:"name"`
	}
)

func NewCreateTagController(u usecases.CreateTagUsecase, p presenters.CreateTagPresenter) CreateTagController {
	return &createTagController{
		usecase:   u,
		presenter: p,
	}
}

func (c *createTagController) Handle(w http.ResponseWriter, r *http.Request) {
	var req createTagRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.Name) < 1 {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	i := &inputs.CreateTagInput{
		Name: req.Name,
	}
	_, err = c.usecase.Execute(r.Context(), i)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
