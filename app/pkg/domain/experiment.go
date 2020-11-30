package domain

type Experiment struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Results []Result `json:"results"`
}
