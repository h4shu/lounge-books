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
		ISBN           string `json:"isbn"`
		Title          string `json:"title"`
		Description    string `json:"description"`
		CoverLink      string `json:"cover_link"`
		PublishedYear  int    `json:"published_year"`
		PublishedMonth int    `json:"published_month"`
		PublishedDay   int    `json:"published_day"`
		Author         string `json:"author"`
		Publisher      string `json:"publisher"`
		TagIDs         []int  `json:"tag_ids"`
		PageCount      int    `json:"page_count"`
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

	publishedAt, err := valueobjects.NewPublishedAt(req.PublishedYear, req.PublishedMonth, req.PublishedDay)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
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
		TagIDs:      req.TagIDs,
	}
	_, err = c.usecase.Execute(r.Context(), i)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
