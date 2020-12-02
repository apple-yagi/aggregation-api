package domain

type Result struct {
	ID           uint     `json:"id"`
	Label        string   `json:"label"`
	Value        []string `json:"value"`
	Unit         string   `json:"unit"`
	ExperimentID uint     `json:"experiment_id"`
}
