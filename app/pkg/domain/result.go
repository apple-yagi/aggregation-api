package domain

type Result struct {
	ID           uint     `json:"id"`
	Label        string   `json:"label"`
	Value        []string `json:"value"`
	ExperimentID uint     `json:"experiment_id"`
}
