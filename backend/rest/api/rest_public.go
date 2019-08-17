package api

import (
	"github.com/go-chi/render"
	"github.com/icheliadinski/cardinal/store"
	"net/http"
)

type public struct {
	dataService pubStore
}

type pubStore interface {
	Collect() ([]store.PageSpeed, error)
}

func (s *public) collectPageSpeedCtrl(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, "{\"score\": 99, \"url\": \"toryburch.com\"}")
}
