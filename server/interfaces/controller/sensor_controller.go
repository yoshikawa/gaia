package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// SensorController model
type SensorController struct {
	Interactor usecase.SensorInteractor
}

// NewSensorController return the SensorController
func NewSensorController(SQLHandler database.SQLHandler) *SensorController {
	return &SensorController{
		Interactor: usecase.SensorInteractor{
			SensorRepository: &database.SensorRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the sensor
func (controller *SensorController) Create(c Context) {
	s := domain.Sensor{}
	c.Bind(&s)
	sensor, err := controller.Interactor.Add(s)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, sensor)
}

// Index is to index the sensors
func (controller *SensorController) Index(c Context) {
	sensors, err := controller.Interactor.Sensors()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, sensors)
}

// Show is to show the sensor
func (controller *SensorController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	sensor, err := controller.Interactor.SensorByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, sensor)
}
