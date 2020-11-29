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
			Label string   `json:"title"`
			Value []string `json:"value"`
		}
		Response struct {
			ResultID int `json:"result_id"`
		}
	)
	req := Request{}
	c.Bind(&req)
	result := domain.Result{Label: req.Label, Value: req.Value}

	id, err := controller.Interactor.Add(result)
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
			ID string `json:"id"`
		}
		Response struct {
			Result domain.Result `json:"result"`
		}
	)
	req := Request{}
	c.Bind(&req)

	r, err := controller.Interactor.FindByID(req.ID)
	if err != nil {
		controller.Interactor.Logger.Log(errors.Wrap(err, "result_controller: not found result"))
		c.JSON(404, NewError(404, err.Error()))
		return
	}
	res := Response{Result: r}
	c.JSON(200, res)
}
