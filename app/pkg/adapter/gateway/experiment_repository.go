package gateway

import (
	"aggregation-mod/pkg/domain"

	"github.com/jinzhu/gorm"
)

type (
	ExperimentRepository struct {
		Conn *gorm.DB
	}

	Experiment struct {
		gorm.Model
		Title    string `gorm:"size:20;not null;unique"`
		Results  []Result
		Interval int
		Count    int
	}
)

func (r *ExperimentRepository) Store(d domain.Experiment) (id int, err error) {
	n := len(d.Results)
	results := make([]Result, n)
	for i := 0; i < n; i++ {
		results[i].ID = d.Results[i].ID
		results[i].Label = d.Results[i].Label
		results[i].Value = d.Results[i].Value
		results[i].Unit = d.Results[i].Unit
		results[i].Color = d.Results[i].Color
		results[i].CreatedAt = d.Results[i].CreatedAt
		results[i].UpdatedAt = d.Results[i].UpdatedAt
	}

	experiment := &Experiment{
		Title:    d.Title,
		Results:  results,
		Interval: d.Interval,
		Count:    d.Count,
	}

	if err = r.Conn.Create(experiment).Error; err != nil {
		return
	}

	return int(experiment.ID), nil
}

func (r *ExperimentRepository) FindByID(id string) (d domain.Experiment, err error) {
	experiment := Experiment{}
	if err = r.Conn.Preload("Results").First(&experiment, id).Error; err != nil {
		return
	}

	n := len(experiment.Results)
	results := make([]domain.Result, n)
	for i := 0; i < n; i++ {
		results[i].ID = experiment.Results[i].ID
		results[i].Label = experiment.Results[i].Label
		results[i].Value = experiment.Results[i].Value
		results[i].Color = experiment.Results[i].Color
		results[i].CreatedAt = experiment.Results[i].CreatedAt
		results[i].UpdatedAt = experiment.Results[i].UpdatedAt
		results[i].ExperimentID = experiment.Results[i].ExperimentID
		results[i].Unit = experiment.Results[i].Unit
	}

	d = domain.Experiment{
		ID:        experiment.ID,
		Title:     experiment.Title,
		Results:   results,
		Interval:  experiment.Interval,
		Count:     experiment.Count,
		CreatedAt: experiment.CreatedAt,
		UpdatedAt: experiment.UpdatedAt,
	}

	return
}

func (r *ExperimentRepository) FindByTitle(title string) (d []domain.Experiment, err error) {
	experiments := []Experiment{}
	if err = r.Conn.Where("title LIKE ?", "%"+title+"%").Find(&experiments).Error; err != nil {
		return
	}

	n := len(experiments)
	d = make([]domain.Experiment, n)
	for i := 0; i < n; i++ {
		d[i].ID = experiments[i].ID
		d[i].Title = experiments[i].Title
		d[i].Count = experiments[i].Count
		d[i].Interval = experiments[i].Interval
		d[i].CreatedAt = experiments[i].CreatedAt
		d[i].UpdatedAt = experiments[i].UpdatedAt
	}
	return
}

func (r *ExperimentRepository) FindAll() (d []domain.Experiment, err error) {
	experiments := []Experiment{}
	if err = r.Conn.Find(&experiments).Error; err != nil {
		return
	}

	n := len(experiments)
	d = make([]domain.Experiment, n)
	for i := 0; i < n; i++ {
		d[i].ID = experiments[i].ID
		d[i].Title = experiments[i].Title
		d[i].Interval = experiments[i].Interval
		d[i].Count = experiments[i].Count
		d[i].CreatedAt = experiments[i].CreatedAt
		d[i].UpdatedAt = experiments[i].UpdatedAt
	}
	return
}

func (r *ExperimentRepository) Update(d domain.Experiment, i string) (id int, err error) {
	n := len(d.Results)
	results := make([]Result, n)
	for i := 0; i < n; i++ {
		results[i].ID = d.Results[i].ID
		results[i].Label = d.Results[i].Label
		results[i].Value = d.Results[i].Value
		results[i].Unit = d.Results[i].Unit
		results[i].Color = d.Results[i].Color
		results[i].CreatedAt = d.Results[i].CreatedAt
		results[i].UpdatedAt = d.Results[i].UpdatedAt
	}

	experiment := Experiment{}
	if err = r.Conn.First(&experiment, i).Error; err != nil {
		return
	}

	experiment.Title = d.Title
	experiment.Interval = d.Interval
	experiment.Count = d.Count
	experiment.Results = results
	if err = r.Conn.Save(experiment).Error; err != nil {
		return
	}

	return int(experiment.ID), nil
}

func (r *ExperimentRepository) Delete(id string) (deleted_id int, err error) {
	experiment := Experiment{}
	if err = r.Conn.Preload("Results").First(&experiment, id).Error; err != nil {
		return
	}

	if err = r.Conn.Delete(&experiment).Error; err != nil {
		return
	}

	return int(experiment.ID), nil
}
