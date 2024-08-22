package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/h4shu/lounge-books/adapter/presenters"
	"github.com/h4shu/lounge-books/application/inputs"
	"github.com/h4shu/lounge-books/application/usecases"
	"github.com/h4shu/lounge-books/domain/valueobjects"
)

type (
	CreateBookController interface {
		Handle(w http.ResponseWriter, r *http.Request)
	}
	createBookControllerImpl struct {
		usecase   usecases.CreateBookUsecase
		presenter presenters.CreateBookPresenter
	}
	createBookRequest struct {
		ISBN        string `json:"isbn"`
		Title       string `json:"title"`
		Description string `json:"description"`
		CoverLink   string `json:"cover_link"`
		PublishedAt string `json:"published_at"`
		Author      string `json:"author"`
		Publisher   string `json:"publisher"`
		PageCount   int    `json:"page_count"`
	}
)

func NewCreateBookController(u usecases.CreateBookUsecase, p presenters.CreateBookPresenter) CreateBookController {
	return &createBookControllerImpl{
		usecase:   u,
		presenter: p,
	}
}

func (c *createBookControllerImpl) Handle(w http.ResponseWriter, r *http.Request) {
	var req createBookRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	publishedAt, err := valueobjects.NewPublishedAtFromStr(req.PublishedAt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	i := &inputs.CreateBookInput{
		ISBN:        valueobjects.NewISBN(req.ISBN),
		Title:       req.Title,
		Description: req.Description,
		CoverLink:   req.CoverLink,
		PublishedAt: publishedAt,
		Author:      valueobjects.NewAuthor(req.Author),
		Publisher:   req.Publisher,
		PageCount:   req.PageCount,
	}
	_, err = c.usecase.Execute(r.Context(), i)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
