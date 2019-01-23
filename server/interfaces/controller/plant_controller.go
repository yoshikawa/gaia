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

// Create is to create the plant
func (controller *PlantController) Create(c Context) {
	p := domain.Plant{}
	c.Bind(&p)
	plant, err := controller.Interactor.Add(p)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, plant)
}

// Index is to index the plants
func (controller *PlantController) Index(c Context) {
	plants, err := controller.Interactor.Plants()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, plants)
}

// Show is to show the plant
func (controller *PlantController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	plant, err := controller.Interactor.PlantByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, plant)
}
