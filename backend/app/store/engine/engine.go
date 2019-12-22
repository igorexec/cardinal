package engine

import (
	"github.com/igorexec/cardinal/app/store"
	"time"
)

type Interface interface {
	Create(pageSpeed store.PageSpeed) (pageSpeedID string, err error)
	Get(from time.Time, to time.Time) ([]store.PageSpeed, error)
}
