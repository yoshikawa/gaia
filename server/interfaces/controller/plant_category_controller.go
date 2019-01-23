package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// PlantCategoryController model
type PlantCategoryController struct {
	Interactor usecase.PlantCategoryInteractor
}

// NewPlantCategoryController return the PlantCategoryController
func NewPlantCategoryController(SQLHandler database.SQLHandler) *PlantCategoryController {
	return &PlantCategoryController{
		Interactor: usecase.PlantCategoryInteractor{
			PlantCategoryRepository: &database.PlantCategoryRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the plant category
func (controller *PlantCategoryController) Create(c Context) {
	pc := domain.PlantCategory{}
	c.Bind(&pc)
	plantCategory, err := controller.Interactor.Add(pc)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, plantCategory)
}

// Index is to index the plant categories
func (controller *PlantCategoryController) Index(c Context) {
	plantCategories, err := controller.Interactor.PlantCategories()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, plantCategories)
}

// Show is to show the plant category
func (controller *PlantCategoryController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	plantCategory, err := controller.Interactor.PlantCategoryByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, plantCategory)
}
