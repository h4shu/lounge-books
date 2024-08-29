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
	ListTagController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}
	listTagController struct {
		usecase   usecases.ListTagUsecase
		presenter presenters.ListTagPresenter
	}
)

func NewListTagController(u usecases.ListTagUsecase, p presenters.ListTagPresenter) ListTagController {
	return &listTagController{
		usecase:   u,
		presenter: p,
	}
}

func (c *listTagController) Handle(w http.ResponseWriter, r *http.Request) {
	i := &inputs.ListTagInput{}
	o, err := c.usecase.Execute(r.Context(), i)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := c.presenter.Output(o)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
