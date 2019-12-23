package collect

import "time"

type PageSpeedCollector struct {
	Token string
}

func NewPageSpeedCollector(token string) *PageSpeedCollector {
	return &PageSpeedCollector{Token: token}
}

func (c *PageSpeedCollector) Do(page string, from time.Time, to time.Time) (int, error) {
	return 0, nil
}
