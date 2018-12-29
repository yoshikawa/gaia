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

// Create is to create the field
func (controller *ObservablePropertyController) Create(c Context) {
	f := domain.ObservableProperty{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *ObservablePropertyController) Index(c Context) {
	fields, err := controller.Interactor.ObservableProperties()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *ObservablePropertyController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.ObservablePropertyByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
