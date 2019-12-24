package api

import (
	"github.com/go-chi/render"
	"github.com/igorexec/cardinal/app/collect"
	"github.com/igorexec/cardinal/app/rest"
	"github.com/igorexec/cardinal/app/store"
	"log"
	"net/http"
)

type private struct {
	dataService privStore

	pageSpeedCollector *collect.PageSpeedCollector
}

type privStore interface {
	pubStore
	Create(pageSpeed store.PageSpeed) (id string, err error)
}

func (s *private) collectCtrl(w http.ResponseWriter, r *http.Request) {
	log.Printf("[info] pagespeed collector initiated")

	pageSpeed := store.PageSpeed{}
	if err := render.DecodeJSON(http.MaxBytesReader(w, r.Body, hardBodyLimit), &pageSpeed); err != nil {
		rest.SendErrorJSON(w, r, http.StatusBadRequest, err, "wrong page name", rest.ErrPageValidation)
		return
	}

	res, err := s.pageSpeedCollector.Do(pageSpeed.Page)
	if err != nil {
		rest.SendErrorJSON(w, r, http.StatusInternalServerError, err, "failed to collect", rest.ErrCollectFail)
		return
	}

	id, err := s.dataService.Create(res)
	if err != nil {
		rest.SendErrorJSON(w, r, http.StatusInternalServerError, err, "failed to save pagespeed data", rest.ErrSaveFail)
	}

	pageSpeed.ID = id

	render.Status(r, http.StatusOK)
	render.JSON(w, r, pageSpeed)
}
