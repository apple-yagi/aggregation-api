package controllers

import (
	"aggregation-mod/pkg/adapter/gateway"
	"aggregation-mod/pkg/adapter/interfaces"
	"aggregation-mod/pkg/domain"
	"aggregation-mod/pkg/usecase"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type ResultController struct {
	Interactor usecase.ResultInteractor
}

func NewResultController(conn *gorm.DB, logger interfaces.Logger) *ResultController {
	return &ResultController{
		Interactor: usecase.ResultInteractor{
			ResultRepository: &gateway.ResultRepository{
				Conn: conn,
			},
			Logger: logger,
		},
	}
}

func (controller *ResultController) Create(c interfaces.Context) {
	type (
		Request struct {
			Label string   `json:"label"`
			Value []string `json:"value"`
			Color string   `json:"color"`
			Unit  string   `json:"unit"`
		}
		Response struct {
			ResultID int `json:"result_id"`
		}
	)
	req := Request{}
	if err := c.Bind(&req); err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: bad request"))
		c.JSON(400, NewError(400, err.Error()))
		return
	}

	e_id := c.Param("experiment_id")

	result := domain.Result{Label: req.Label, Value: req.Value, Color: req.Color, Unit: req.Unit}

	id, err := controller.Interactor.Add(result, e_id)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: cannot add result"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{ResultID: id}
	c.JSON(201, res)
}

func (controller *ResultController) Show(c interfaces.Context) {
	type (
		Request struct {
			ID string
		}
		Response struct {
			ID           uint     `json:"id"`
			Label        string   `json:"label"`
			Value        []string `json:"value"`
			Unit         string   `json:"unit"`
			Color        string   `json:"color"`
			ExperimentID uint     `json:"experiment_id"`
		}
	)
	req := Request{}
	req.ID = c.Param("id")

	r, err := controller.Interactor.FindByID(req.ID)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: not found result"))
		c.JSON(404, NewError(404, err.Error()))
		return
	}
	res := Response{ID: r.ID, Label: r.Label, Value: r.Value, Unit: r.Unit, ExperimentID: r.ExperimentID}
	c.JSON(200, res)
}

func (controller *ResultController) Index(c interfaces.Context) {
	r, err := controller.Interactor.FindAll()
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: findall error"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	c.JSON(200, r)
}

func (controller *ResultController) Delete(c interfaces.Context) {
	type (
		Request struct {
			ID string
		}
		Response struct {
			ResultID int `json:"result_id"`
		}
	)
	req := Request{}
	req.ID = c.Param("id")

	r, err := controller.Interactor.Remove(req.ID)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: cannot remove result"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}
	res := Response{ResultID: r}
	c.JSON(200, res.ResultID)
}

func (controller *ResultController) Update(c interfaces.Context) {
	type (
		Request struct {
			Label string   `json:"label"`
			Value []string `json:"value"`
			Unit  string   `json:"unit"`
			Color string   `json:"color"`
		}
		Response struct {
			ResultID int `json:"result_id"`
		}
	)
	req := Request{}
	var err error

	err = c.Bind(&req)
	i := c.Param("id")

	if err != nil || i == "" {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: bad request"))
		c.JSON(400, NewError(400, "Bad Request"))
		return
	}
	result := domain.Result{Label: req.Label, Value: req.Value}

	id, err := controller.Interactor.Update(result, i)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: cannot update result"))
		c.JSON(500, NewError(500, err.Error()))
		return
	}

	res := Response{ResultID: id}
	c.JSON(200, res.ResultID)
}
