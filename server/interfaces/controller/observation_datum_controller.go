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

// Create is to create the observation dataum
func (controller *ObservationDatumController) Create(c Context) {
	od := domain.ObservationDatum{}
	c.Bind(&od)
	observationDatum, err := controller.Interactor.Add(od)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, observationDatum)
}

// Index is to index the observation data
func (controller *ObservationDatumController) Index(c Context) {
	observationData, err := controller.Interactor.ObservationData()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, observationData)
}

// Show is to show the observation dataum
func (controller *ObservationDatumController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	observationDataum, err := controller.Interactor.ObservationDatumByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, observationDataum)
}
