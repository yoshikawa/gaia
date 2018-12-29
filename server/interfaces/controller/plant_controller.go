package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// PlantController model
type PlantController struct {
	Interactor usecase.PlantInteractor
}

// NewPlantController return the PlantController
func NewPlantController(SQLHandler database.SQLHandler) *PlantController {
	return &PlantController{
		Interactor: usecase.PlantInteractor{
			PlantRepository: &database.PlantRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *PlantController) Create(c Context) {
	f := domain.Plant{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *PlantController) Index(c Context) {
	fields, err := controller.Interactor.Plants()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *PlantController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.PlantByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
