package usecase

import (
	"aggregation-mod/pkg/domain"
	"aggregation-mod/pkg/usecase/interfaces"
)

type ExperimentInteractor struct {
	ExperimentRepository interfaces.ExperimentRepository
	Logger               interfaces.Logger
}

func (i *ExperimentInteractor) Add(e domain.Experiment) (int, error) {
	i.Logger.Log("store experiment!")
	return i.ExperimentRepository.Store(e)
}

func (i *ExperimentInteractor) Remove(id string) (int, error) {
	i.Logger.Log("remove experiment!")
	return i.ExperimentRepository.Delete(id)
}

func (i *ExperimentInteractor) FindAll() ([]domain.Experiment, error) {
	i.Logger.Log("findall experiment")
	return i.ExperimentRepository.FindAll()
}

func (i *ExperimentInteractor) FindByID(id string) (domain.Experiment, error) {
	i.Logger.Log("find by id experiment")
	return i.ExperimentRepository.FindByID(id)
}

func (i *ExperimentInteractor) FindByTitle(title string) ([]domain.Experiment, error) {
	i.Logger.Log("find by title experiment")
	return i.ExperimentRepository.FindByTitle(title)
}
