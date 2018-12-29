package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// ObservationPositionController model
type ObservationPositionController struct {
	Interactor usecase.ObservationPositionInteractor
}

// NewObservationPositionController return the ObservationPositionController
func NewObservationPositionController(SQLHandler database.SQLHandler) *ObservationPositionController {
	return &ObservationPositionController{
		Interactor: usecase.ObservationPositionInteractor{
			ObservationPositionRepository: &database.ObservationPositionRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *ObservationPositionController) Create(c Context) {
	f := domain.ObservationPosition{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *ObservationPositionController) Index(c Context) {
	fields, err := controller.Interactor.ObservationPositions()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *ObservationPositionController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.ObservationPositionByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
