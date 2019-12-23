package collect

import "time"

type Collect interface {
	Do(page string, from time.Time, to time.Time) (int, error)
}
