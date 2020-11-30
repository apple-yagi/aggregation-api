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
		Value        pq.StringArray `gorm:"type:text[]"`
		ExperimentID uint
	}
)

func (r *ResultRepository) Store(d domain.Result) (id int, err error) {
	result := &Result{
		Label: d.Label,
		// Value: d.Value,
	}

	if err = r.Conn.Create(result).Error; err != nil {
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
		ID:    result.ID,
		Label: result.Label,
		// Value: result.Value,
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
		// d[i].Value = results[i].Value
	}

	return
}
