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
	ListBookController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}
	listBookControllerImpl struct {
		usecase   usecases.ListBookUsecase
		presenter presenters.ListBookPresenter
	}
)

func NewListBookController(u usecases.ListBookUsecase, p presenters.ListBookPresenter) ListBookController {
	return &listBookControllerImpl{
		usecase:   u,
		presenter: p,
	}
}

func (c *listBookControllerImpl) Handle(w http.ResponseWriter, r *http.Request) {
	i := &inputs.ListBookInput{}
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
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
