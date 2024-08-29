package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/h4shu/lounge-books/adapter/presenters"
	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/usecases"
	"github.com/h4shu/lounge-books/domain/entities"
)

type (
	GetTagController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}
	getTagController struct {
		usecase   usecases.GetTagUsecase
		presenter presenters.GetTagPresenter
	}
)

func NewGetTagController(u usecases.GetTagUsecase, p presenters.GetTagPresenter) GetTagController {
	return &getTagController{
		usecase:   u,
		presenter: p,
	}
}

func (c *getTagController) Handle(w http.ResponseWriter, r *http.Request) {
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

	i := &inputs.GetTagInput{
		ID: id,
	}
	o, err := c.usecase.Execute(r.Context(), i)
	if err != nil {
		if errors.Is(err, entities.ErrTagNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}

	res := c.presenter.Output(o)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
