package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/h4shu/lounge-books/adapter/presenters"
	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/usecases"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type (
	GetBookController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}
	getBookControllerImpl struct {
		usecase   usecases.GetBookUsecase
		presenter presenters.GetBookPresenter
	}
)

func NewGetBookController(u usecases.GetBookUsecase, p presenters.GetBookPresenter) GetBookController {
	return &getBookControllerImpl{
		usecase:   u,
		presenter: p,
	}
}

func (c *getBookControllerImpl) Handle(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p[len(p)-1])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i := &inputs.GetBookInput{
		ID: valueobjects.NewBookID(id),
	}
	o, err := c.usecase.Execute(r.Context(), i)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if o.Book == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res := c.presenter.Output(o)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
