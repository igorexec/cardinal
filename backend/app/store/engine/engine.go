package engine

import (
	"context"
	"github.com/igorexec/cardinal/app/store"
	"time"
)

const database = "cardinal"

type Interface interface {
	Create(context.Context, store.PageSpeed) (pageSpeedID string, err error)
	Get(ctx context.Context, from time.Time, to time.Time) ([]store.PageSpeed, error)
	Close(context.Context) error
}
