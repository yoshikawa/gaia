package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// ObservationDatumController model
type ObservationDatumController struct {
	Interactor usecase.ObservationDatumInteractor
}

// NewObservationDatumController return the ObservationDatumController
func NewObservationDatumController(SQLHandler database.SQLHandler) *ObservationDatumController {
	return &ObservationDatumController{
		Interactor: usecase.ObservationDatumInteractor{
			ObservationDatumRepository: &database.ObservationDatumRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *ObservationDatumController) Create(c Context) {
	f := domain.ObservationDatum{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *ObservationDatumController) Index(c Context) {
	fields, err := controller.Interactor.ObservationData()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *ObservationDatumController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.ObservationDatumByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
