package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// SensingDeviceStatusController model
type SensingDeviceStatusController struct {
	Interactor usecase.SensingDeviceStatusInteractor
}

// NewSensingDeviceStatusController return the SensingDeviceStatusController
func NewSensingDeviceStatusController(SQLHandler database.SQLHandler) *SensingDeviceStatusController {
	return &SensingDeviceStatusController{
		Interactor: usecase.SensingDeviceStatusInteractor{
			SensingDeviceStatusRepository: &database.SensingDeviceStatusRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *SensingDeviceStatusController) Create(c Context) {
	f := domain.SensingDeviceStatus{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *SensingDeviceStatusController) Index(c Context) {
	fields, err := controller.Interactor.SensingDeviceStatuses()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *SensingDeviceStatusController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.SensingDeviceStatusByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
