package collect

type Collect interface {
	Do(page string) (int, error)
}
