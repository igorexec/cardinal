package service

import (
	"github.com/icheliadinski/cardinal/store"
	"github.com/icheliadinski/cardinal/store/engine"
	"github.com/pkg/errors"
)

type DataStore struct {
	Engine engine.Interface
}

func (s *DataStore) Save(pageSpeed store.PageSpeed) error {
	ps, err := s.preparePageSpeed(pageSpeed)
	if err != nil {
		return errors.Wrap(err, "failed to prepare page speed")
	}
	return s.Engine.Save(ps)
}

func (s *DataStore) preparePageSpeed(pageSpeed store.PageSpeed) (store.PageSpeed, error) {
	if pageSpeed.Score == 0 {
		return store.PageSpeed{}, errors.Errorf("page speed index is 0 for %s", pageSpeed.Page)
	}
	return pageSpeed, nil
}
