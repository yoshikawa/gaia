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

// Create is to create the observation position
func (controller *ObservationPositionController) Create(c Context) {
	op := domain.ObservationPosition{}
	c.Bind(&op)
	observationPosition, err := controller.Interactor.Add(op)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, observationPosition)
}

// Index is to index the observation positions
func (controller *ObservationPositionController) Index(c Context) {
	observationPositions, err := controller.Interactor.ObservationPositions()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, observationPositions)
}

// Show is to show the observation position
func (controller *ObservationPositionController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	observationPosition, err := controller.Interactor.ObservationPositionByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, observationPosition)
}
