package usecase

import (
	"aggregation-mod/pkg/domain"
	"aggregation-mod/pkg/usecase/interfaces"
)

type ResultInteractor struct {
	ResultRepository interfaces.ResultRepository
	Logger           interfaces.Logger
}

func (i *ResultInteractor) Add(r domain.Result) (int, error) {
	i.Logger.Log("store result!")
	return i.ResultRepository.Store(r)
}

func (i *ResultInteractor) FindAll() ([]domain.Result, error) {
	i.Logger.Log("find all result")
	return i.ResultRepository.FindAll()
}

func (i *ResultInteractor) FindByID(id string) (domain.Result, error) {
	i.Logger.Log("find by id result")
	return i.ResultRepository.FindByID(id)
}
