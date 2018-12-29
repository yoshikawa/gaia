package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// GatewayDeviceController model
type GatewayDeviceController struct {
	Interactor usecase.GatewayDeviceInteractor
}

// NewGatewayDeviceController return the GatewayDeviceController
func NewGatewayDeviceController(SQLHandler database.SQLHandler) *GatewayDeviceController {
	return &GatewayDeviceController{
		Interactor: usecase.GatewayDeviceInteractor{
			GatewayDeviceRepository: &database.GatewayDeviceRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *GatewayDeviceController) Create(c Context) {
	f := domain.GatewayDevice{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *GatewayDeviceController) Index(c Context) {
	fields, err := controller.Interactor.GatewayDevices()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *GatewayDeviceController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.GatewayDeviceByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
