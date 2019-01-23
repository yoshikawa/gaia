package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// SensingDeviceController model
type SensingDeviceController struct {
	Interactor usecase.SensingDeviceInteractor
}

// NewSensingDeviceController return the SensingDeviceController
func NewSensingDeviceController(SQLHandler database.SQLHandler) *SensingDeviceController {
	return &SensingDeviceController{
		Interactor: usecase.SensingDeviceInteractor{
			SensingDeviceRepository: &database.SensingDeviceRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the sensing device
func (controller *SensingDeviceController) Create(c Context) {
	sd := domain.SensingDevice{}
	c.Bind(&sd)
	sensingDevice, err := controller.Interactor.Add(sd)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, sensingDevice)
}

// Index is to index the sensing devices
func (controller *SensingDeviceController) Index(c Context) {
	sensingDevices, err := controller.Interactor.SensingDevices()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, sensingDevices)
}

// Show is to show the sensing device
func (controller *SensingDeviceController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	sensingDevice, err := controller.Interactor.SensingDeviceByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, sensingDevice)
}
