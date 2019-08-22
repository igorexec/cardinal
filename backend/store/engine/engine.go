package engine

import (
	"github.com/icheliadinski/cardinal/store"
	"time"
)

type Interface interface {
	Save(pageSpeed store.PageSpeed) error
	Get(from time.Time, to time.Time) ([]store.PageSpeed, error)
}
