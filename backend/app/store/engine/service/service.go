package service

import (
	"context"
	"github.com/igorexec/cardinal/app/store"
	"github.com/igorexec/cardinal/app/store/engine"
	"time"
)

type DataStore struct {
	Engine engine.Interface
}

func (d *DataStore) Create(speed store.PageSpeed) (id string, err error) {
	ctx := context.Background()
	return d.Engine.Create(ctx, speed)
}

func (d *DataStore) Get(page string, from time.Time, to time.Time) ([]store.PageSpeed, error) {
	ctx := context.Background()
	return d.Engine.Get(ctx, page, from, to)
}

func (d *DataStore) Close() error {
	ctx := context.Background()
	return d.Engine.Close(ctx)
}
