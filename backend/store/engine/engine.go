package engine

import "github.com/icheliadinski/cardinal/store"

type Interface interface {
	Save(pageSpeed store.PageSpeed) error
}
