package controllers

import (
	"aggregation-mod/pkg/adapter/gateway"
	"aggregation-mod/pkg/adapter/interfaces"
	"aggregation-mod/pkg/domain"
	"aggregation-mod/pkg/usecase"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ExperimentController struct {
	Interactor usecase.ExperimentInteractor
}

func NewExperimentController(conn *gorm.DB, logger interfaces.Logger) *ExperimentController {
	return &ExperimentController{
		Interactor: usecase.ExperimentInteractor{
			ExperimentRepository: &gateway.ExperimentRepository{
				Conn: conn,
			},
			Logger: logger,
		},
	}
}

func (controller *ExperimentController) Create(c interfaces.Context) {
	type (
		Request struct {
			Title   string          `json:"title"`
			Results []domain.Result `json:"results"`
		}
		Response struct {
			ExperimentID int `json:"experiment_id"`
		}
	)
	req := Request{}
	if err := c.Bind(&req); err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "experiment_controller: bad request"))
		c.JSON(400, NewError(400, err.Error()))
		return
	}
	experiment := domain.Experiment{Title: req.Title, Results: req.Results}

	id, err := controller.Interactor.Add(experiment)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "experiment_controller: cannot add experiment"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{ExperimentID: id}
	c.JSON(201, res)
}

func (controller *ExperimentController) Show(c interfaces.Context) {
	type (
		Request struct {
			ID string
		}
		Response struct {
			Experiment domain.Experiment `json:"experiment"`
		}
	)
	req := Request{}
	req.ID = c.Param("id")

	r, err := controller.Interactor.FindByID(req.ID)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "experiment_controller: not found experiment"))
		c.JSON(404, NewError(404, err.Error()))
		return
	}
	res := Response{Experiment: r}
	c.JSON(200, res)
}

func (controller *ExperimentController) Index(c interfaces.Context) {
	type (
		Response struct {
			Experiments []domain.Experiment `json:"experiments"`
		}
	)

	r, err := controller.Interactor.FindAll()
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: findall error"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{Experiments: r}
	c.JSON(200, res)
}
