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
		Title   string          `gorm:"size:20;not null"`
		Results []domain.Result `gorm:"foreignKey:ExperimentRefer"`
	}
)

func (r *ExperimentRepository) Store(d domain.Experiment) (id int, err error) {
	experiment := &Experiment{
		Title: d.Title,
	}

	if err = r.Conn.Create(experiment).Error; err != nil {
		return
	}

	return int(experiment.ID), nil
}

func (r *ExperimentRepository) FindByID(id string) (d domain.Experiment, err error) {
	experiment := Experiment{}
	if err = r.Conn.First(&experiment, id).Error; err != nil {
		return
	}

	d = domain.Experiment{
		Title:   experiment.Title,
		Results: experiment.Results,
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
		d[i].ID = int(experiments[i].ID)
		d[i].Title = experiments[i].Title
		d[i].Results = experiments[i].Results
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
		d[i].ID = int(experiments[i].ID)
		d[i].Title = experiments[i].Title
		d[i].Results = experiments[i].Results
	}
	return
}
