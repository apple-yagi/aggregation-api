package usecase

import (
	"aggregation-mod/pkg/domain"
	"aggregation-mod/pkg/usecase/interfaces"
)

type ResultInteractor struct {
	ResultRepository interfaces.ResultRepository
	Logger           interfaces.Logger
}

func (i *ResultInteractor) Add(r domain.Result, experiment_id string) (int, error) {
	i.Logger.Log("store result!")
	return i.ResultRepository.Store(r, experiment_id)
}

func (i *ResultInteractor) Remove(id string) (int, error) {
	i.Logger.Log("remove result!")
	return i.ResultRepository.Delete(id)
}

func (i *ResultInteractor) Update(r domain.Result, id string) (int, error) {
	i.Logger.Log("update result!")
	return i.ResultRepository.Update(r, id)
}

func (i *ResultInteractor) FindAll() ([]domain.Result, error) {
	i.Logger.Log("find all result")
	return i.ResultRepository.FindAll()
}

func (i *ResultInteractor) FindByID(id string) (domain.Result, error) {
	i.Logger.Log("find by id result")
	return i.ResultRepository.FindByID(id)
}
