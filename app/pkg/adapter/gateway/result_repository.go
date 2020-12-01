package gateway

import (
	"aggregation-mod/pkg/domain"

	"github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

type (
	ResultRepository struct {
		Conn *gorm.DB
	}

	Result struct {
		gorm.Model
		Label        string         `gorm:"size:20;not null"`
		Value        pq.StringArray `gorm:"type:text[];not null"`
		ExperimentID uint
	}
)

func (r *ResultRepository) Store(d domain.Result, experiment_id string) (id int, err error) {
	experiment := Experiment{}
	if err = r.Conn.First(&experiment, experiment_id).Error; err != nil {
		return
	}

	result := &Result{
		Label:        d.Label,
		Value:        d.Value,
		ExperimentID: experiment.ID,
	}

	if err = r.Conn.Create(result).Error; err != nil {
		return
	}

	return int(result.ID), nil
}

func (r *ResultRepository) Update(d domain.Result, i string) (id int, err error) {
	result := Result{}
	if err = r.Conn.First(&result, i).Error; err != nil {
		return
	}

	result.Label = d.Label
	result.Value = d.Value

	if err = r.Conn.Update(&result).Error; err != nil {
		return
	}

	return int(result.ID), nil
}

func (r *ResultRepository) FindByID(id string) (d domain.Result, err error) {
	result := Result{}
	if err = r.Conn.First(&result, id).Error; err != nil {
		return
	}

	d = domain.Result{
		ID:           result.ID,
		Label:        result.Label,
		Value:        result.Value,
		ExperimentID: result.ExperimentID,
	}

	return
}

func (r *ResultRepository) FindAll() (d []domain.Result, err error) {
	results := []Result{}
	if err = r.Conn.Find(&results).Error; err != nil {
		return
	}

	n := len(results)
	d = make([]domain.Result, n)
	for i := 0; i < n; i++ {
		d[i].ID = results[i].ID
		d[i].Label = results[i].Label
		d[i].Value = results[i].Value
		d[i].ExperimentID = results[i].ExperimentID
	}

	return
}

func (r *ResultRepository) Delete(id string) (deleted_id int, err error) {
	result := Result{}
	if err = r.Conn.First(&result, id).Error; err != nil {
		return
	}

	if err = r.Conn.Delete(&result).Error; err != nil {
		return
	}

	return int(result.ID), nil
}
