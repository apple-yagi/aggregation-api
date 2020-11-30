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
		Title   string `gorm:"size:20;not null"`
		Results []Result
	}
)

func (r *ExperimentRepository) Store(d domain.Experiment) (id int, err error) {
	n := len(d.Results)
	results := make([]Result, n)
	for i := 0; i < n; i++ {
		results[i].ID = d.Results[i].ID
		results[i].Label = d.Results[i].Label
		results[i].Value = d.Results[i].Value
	}

	experiment := &Experiment{
		Title:   d.Title,
		Results: results,
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
		results[i].ExperimentID = experiment.Results[i].ExperimentID
	}

	d = domain.Experiment{
		ID:      experiment.ID,
		Title:   experiment.Title,
		Results: results,
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
	}
	return
}
