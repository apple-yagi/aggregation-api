package domain

type Experiment struct {
	ID       uint     `json:"id"`
	Title    string   `json:"title"`
	Results  []Result `json:"results"`
	TimeAxis []string `json:"time_axis"`
}
