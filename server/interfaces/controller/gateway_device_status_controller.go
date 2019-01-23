package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// GatewayDeviceStatusController model
type GatewayDeviceStatusController struct {
	Interactor usecase.GatewayDeviceStatusInteractor
}

// NewGatewayDeviceStatusController return the GatewayDeviceStatusController
func NewGatewayDeviceStatusController(SQLHandler database.SQLHandler) *GatewayDeviceStatusController {
	return &GatewayDeviceStatusController{
		Interactor: usecase.GatewayDeviceStatusInteractor{
			GatewayDeviceStatusRepository: &database.GatewayDeviceStatusRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the gateway device status
func (controller *GatewayDeviceStatusController) Create(c Context) {
	gds := domain.GatewayDeviceStatus{}
	c.Bind(&gds)
	gatewayDeviceStatus, err := controller.Interactor.Add(gds)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, gatewayDeviceStatus)
}

// Index is to index the gateway device statuses
func (controller *GatewayDeviceStatusController) Index(c Context) {
	gatewayDeviceStatuses, err := controller.Interactor.GatewayDeviceStatuses()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gatewayDeviceStatuses)
}

// Show is to show the gateway device status by id
func (controller *GatewayDeviceStatusController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	gatewayDeviceStatus, err := controller.Interactor.GatewayDeviceStatusByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gatewayDeviceStatus)
}
