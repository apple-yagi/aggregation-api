package gateway

import (
	"aggregation-mod/pkg/domain"
	"fmt"

	"github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

type (
	ResultRepository struct {
		Conn *gorm.DB
	}

	Result struct {
		gorm.Model
		Label        string `gorm:"size:20;not null"`
		Value        pq.Float64Array
		ExperimentID uint
		Unit         string `gorm:"not null"`
		Color        string
	}
)

func (r *ResultRepository) Store(d domain.Result, experiment_id string) (id int, err error) {
	experiment := Experiment{}
	if err = r.Conn.First(&experiment, experiment_id).Error; err != nil {
		return
	}

	fmt.Println(d)

	result := &Result{
		Label:        d.Label,
		Value:        d.Value,
		ExperimentID: experiment.ID,
		Unit:         d.Unit,
		Color:        d.Color,
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
	result.Unit = d.Unit
	result.Color = d.Color

	if err = r.Conn.Save(&result).Error; err != nil {
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
		Unit:         result.Unit,
		Color:        result.Color,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
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
		d[i].Unit = results[i].Unit
		d[i].Color = results[i].Color
		d[i].CreatedAt = results[i].CreatedAt
		d[i].UpdatedAt = results[i].UpdatedAt
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
