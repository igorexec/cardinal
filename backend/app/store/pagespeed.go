package store

import "time"

type PageSpeed struct {
	ID    string    `json:"id,omitempty" bson:"_id,omitempty"`
	Score int       `json:"score"`
	Page  string    `json:"page"`
	Date  time.Time `json:"time"`
}
