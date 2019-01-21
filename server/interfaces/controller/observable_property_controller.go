package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// ObservablePropertyController model
type ObservablePropertyController struct {
	Interactor usecase.ObservablePropertyInteractor
}

// NewObservablePropertyController return the ObservablePropertyController
func NewObservablePropertyController(SQLHandler database.SQLHandler) *ObservablePropertyController {
	return &ObservablePropertyController{
		Interactor: usecase.ObservablePropertyInteractor{
			ObservablePropertyRepository: &database.ObservablePropertyRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the observableProperty
func (controller *ObservablePropertyController) Create(c Context) {
	op := domain.ObservableProperty{}
	c.Bind(&op)
	observableProperty, err := controller.Interactor.Add(op)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, observableProperty)
}

// Index is to index the observableProperties
func (controller *ObservablePropertyController) Index(c Context) {
	observableProperties, err := controller.Interactor.ObservableProperties()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, observableProperties)
}

// Show is to show the observableProperty
func (controller *ObservablePropertyController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	observableProperty, err := controller.Interactor.ObservablePropertyByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, observableProperty)
}
