package store

import "time"

type PageSpeed struct {
	Page  string    `json:"page"`
	Score int       `json:"score"`
	Date  time.Time `json:"date"`
}
