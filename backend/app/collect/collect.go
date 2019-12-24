package collect

import "github.com/igorexec/cardinal/app/store"

type Collect interface {
	Do(page string) (store.PageSpeed, error)
}
