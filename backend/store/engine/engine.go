package engine

import "github.com/icheliadinski/cardinal/store"

type Interface interface {
}

type Accessor interface {
	Create(pageSpeed store.PageSpeed) (pageSpeedID string, err error) // create new pagespeed data
}
