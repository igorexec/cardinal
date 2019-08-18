package api

import (
	"github.com/go-chi/render"
	"github.com/go-pkgz/lgr"
	"github.com/icheliadinski/cardinal/collector"
	"github.com/icheliadinski/cardinal/rest"
	"github.com/icheliadinski/cardinal/store"
	"net/http"
	"time"
)

type public struct {
	dataService        pubStore
	pageSpeedCollector *collector.PageSpeed
}

type pubStore interface {
	Create(pageSpeed store.PageSpeed) (pageSpeedID string, err error)
	FindSince(since time.Time, to time.Time, site string) ([]store.PageSpeed, error)
}

func (s *public) collectPageSpeedCtrl(w http.ResponseWriter, r *http.Request) {
	lgr.Printf("[INFO] pagespeed collector initiated")

	pageSpeed := store.PageSpeed{}
	if err := render.DecodeJSON(http.MaxBytesReader(w, r.Body, hardBodyLimit), &pageSpeed); err != nil {
		rest.SendErrorJSON(w, r, http.StatusBadRequest, err, "wrong page name", rest.ErrPageValidation)
		return
	}

	res, err := s.pageSpeedCollector.Collect(pageSpeed.Page)
	if err != nil {
		rest.SendErrorJSON(w, r, http.StatusInternalServerError, err, "failed to collect", rest.ErrCollectFail)
		return
	}

	render.JSON(w, r, res)
}
