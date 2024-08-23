package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/h4shu/lounge-books/adapter/presenters"
	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/usecases"
	"github.com/h4shu/lounge-books/domain/entities"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type (
	DeleteBookController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}
	deleteBookControllerImpl struct {
		usecase   usecases.DeleteBookUsecase
		presenter presenters.DeleteBookPresenter
	}
)

func NewDeleteBookController(u usecases.DeleteBookUsecase, p presenters.DeleteBookPresenter) DeleteBookController {
	return &deleteBookControllerImpl{
		usecase:   u,
		presenter: p,
	}
}

func (c *deleteBookControllerImpl) Handle(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(p[len(p)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	i := &inputs.DeleteBookInput{
		ID: valueobjects.NewBookID(id),
	}
	_, err = c.usecase.Execute(r.Context(), i)
	if err != nil {
		if errors.Is(err, entities.ErrBookNotFound) || errors.Is(err, entities.ErrBookAlreadyDeleted) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
