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

// Create is to create the field
func (controller *SensingDeviceController) Create(c Context) {
	f := domain.SensingDevice{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *SensingDeviceController) Index(c Context) {
	fields, err := controller.Interactor.SensingDevices()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *SensingDeviceController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.SensingDeviceByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
