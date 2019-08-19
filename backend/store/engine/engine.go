package engine

import "github.com/icheliadinski/cardinal/store"

type Interface interface {
}

type Accessor interface {
	Save(pageSpeed store.PageSpeed) error // create new pagespeed data
}
