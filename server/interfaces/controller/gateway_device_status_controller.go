package controller

import (
	"strconv"

	"github.com/Pluslab/fieldsensing/server/domain"
	"github.com/Pluslab/fieldsensing/server/interfaces/database"
	"github.com/Pluslab/fieldsensing/server/usecase"
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

// Create is to create the field
func (controller *GatewayDeviceStatusController) Create(c Context) {
	f := domain.GatewayDeviceStatus{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *GatewayDeviceStatusController) Index(c Context) {
	fields, err := controller.Interactor.GatewayDeviceStatuses()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *GatewayDeviceStatusController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.GatewayDeviceStatusByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
