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
	dataService pubStore
	collector   *collector.Collector
}

type pubStore interface {
	Save(pageSpeed store.PageSpeed) error
	Get(from time.Time, to time.Time) ([]store.PageSpeed, error)
}

func (s *public) collectPageSpeedCtrl(w http.ResponseWriter, r *http.Request) {
	lgr.Printf("[INFO] pagespeed collector initiated")

	pageSpeed := store.PageSpeed{}
	if err := render.DecodeJSON(http.MaxBytesReader(w, r.Body, hardBodyLimit), &pageSpeed); err != nil {
		rest.SendErrorJSON(w, r, http.StatusBadRequest, err, "wrong page name", rest.ErrPageValidation)
		return
	}

	res, err := s.collector.CollectPageSpeed(pageSpeed.Page)
	if err != nil {
		rest.SendErrorJSON(w, r, http.StatusInternalServerError, err, "failed to collect", rest.ErrCollectFail)
		return
	}

	if err := s.dataService.Save(res); err != nil {
		rest.SendErrorJSON(w, r, http.StatusInternalServerError, err, "failed to save pagespeed data", rest.ErrSaveFail)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
}

func (s *public) findPageSpeed(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Query().Get("from")
	t := r.URL.Query().Get("to")
	now := time.Now()

	from, err := time.Parse("2006-01-02", f)
	if err != nil {
		from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	}

	to, err := time.Parse("2006-01-02", t)
	if err != nil {
		to = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24)
	}

	ps, err := s.dataService.Get(from, to)
	if err != nil {
		rest.SendErrorJSON(w, r, http.StatusInternalServerError, err, "failed to get pagespeed data", rest.ErrNoData)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, ps)
}
