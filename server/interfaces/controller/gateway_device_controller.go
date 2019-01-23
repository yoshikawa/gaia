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

// Create is to create the gateway device
func (controller *GatewayDeviceController) Create(c Context) {
	gd := domain.GatewayDevice{}
	c.Bind(&gd)
	gatewayDevice, err := controller.Interactor.Add(gd)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, gatewayDevice)
}

// Index is to index the gateway devices
func (controller *GatewayDeviceController) Index(c Context) {
	gatewayDevices, err := controller.Interactor.GatewayDevices()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gatewayDevices)
}

// Show is to show the gateway device
func (controller *GatewayDeviceController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	gatewayDevice, err := controller.Interactor.GatewayDeviceByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gatewayDevice)
}
