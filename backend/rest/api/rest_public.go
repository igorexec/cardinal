package api

import (
	"github.com/icheliadinski/cardinal/store"
	"net/http"
	"time"
)

type public struct {
	dataService pubStore
}

type pubStore interface {
	Collect() error
	FindSince(since time.Time, to time.Time) ([]store.PageSpeed, error)
}

func (s *public) collectPageSpeedCtrl(w http.ResponseWriter, r *http.Request) {
	// TODO: add collecting
}

func (s *public) findSince(w http.ResponseWriter, r *http.Request) {
	// TODO: add logic for since
}
