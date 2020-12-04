package domain

import (
	"time"
)

type Result struct {
	ID           uint      `json:"id"`
	Label        string    `json:"label"`
	Value        []float64 `json:"value"`
	Unit         string    `json:"unit"`
	Color        string    `json:"color"`
	ExperimentID uint      `json:"experiment_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
