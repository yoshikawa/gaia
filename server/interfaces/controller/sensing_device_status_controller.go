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

// Create is to create the sensing device status
func (controller *SensingDeviceStatusController) Create(c Context) {
	sds := domain.SensingDeviceStatus{}
	c.Bind(&sds)
	sensingDeviceStatus, err := controller.Interactor.Add(sds)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, sensingDeviceStatus)
}

// Index is to index the sensing device statuses
func (controller *SensingDeviceStatusController) Index(c Context) {
	sensingDeviceStatuses, err := controller.Interactor.SensingDeviceStatuses()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, sensingDeviceStatuses)
}

// Show is to show the sensing device status
func (controller *SensingDeviceStatusController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	sensingDeviceStatus, err := controller.Interactor.SensingDeviceStatusByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, sensingDeviceStatus)
}
