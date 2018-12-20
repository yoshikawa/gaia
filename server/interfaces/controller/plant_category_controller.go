package controller

import (
	"strconv"

	"github.com/Pluslab/fieldsensing/server/domain"
	"github.com/Pluslab/fieldsensing/server/interfaces/database"
	"github.com/Pluslab/fieldsensing/server/usecase"
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

// Create is to create the field
func (controller *PlantCategoryController) Create(c Context) {
	f := domain.PlantCategory{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *PlantCategoryController) Index(c Context) {
	fields, err := controller.Interactor.PlantCategories()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *PlantCategoryController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.PlantCategoryByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
