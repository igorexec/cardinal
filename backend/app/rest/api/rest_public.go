package api

import (
	"github.com/go-chi/render"
	"github.com/igorexec/cardinal/app/rest"
	"github.com/igorexec/cardinal/app/store"
	"net/http"
	"time"
)

type public struct {
	dataService pubStore
}

type pubStore interface {
	Get(page string, from time.Time, to time.Time) ([]store.PageSpeed, error)
}

func (s *public) findPageSpeed(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Query().Get("from")
	t := r.URL.Query().Get("to")
	page := r.URL.Query().Get("page")
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	from, err := time.Parse("2006-01-02", f)
	if err != nil {
		from = today
	}

	to, err := time.Parse("2006-01-02", t)
	if err != nil {
		to = today.Add(24 * time.Hour)
	}

	ps, err := s.dataService.Get(page, from, to)
	if err != nil {
		rest.SendErrorJSON(w, r, http.StatusInternalServerError, err, "failed to get pagespeed data", rest.ErrNoData)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, ps)
}
