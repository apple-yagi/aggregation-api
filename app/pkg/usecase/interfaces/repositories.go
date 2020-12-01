package interfaces

import (
	"aggregation-mod/pkg/domain"
)

type ResultRepository interface {
	Store(domain.Result, string) (int, error)
	FindByID(string) (domain.Result, error)
	FindAll() ([]domain.Result, error)
	Delete(string) (int, error)
}

type ExperimentRepository interface {
	Store(domain.Experiment) (int, error)
	FindByID(string) (domain.Experiment, error)
	FindByTitle(string) ([]domain.Experiment, error)
	FindAll() ([]domain.Experiment, error)
	Delete(string) (int, error)
}
