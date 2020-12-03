package domain

import (
	"time"
)

type Experiment struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Results   []Result  `json:"results"`
	Interval  int       `json:"interval"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
