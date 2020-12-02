package domain

import (
	"time"
)

type Experiment struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Results   []Result  `json:"results"`
	TimeAxis  []string  `json:"time_axis"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
