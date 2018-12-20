package controller

import (
	"strconv"

	"github.com/Pluslab/fieldsensing/server/domain"
	"github.com/Pluslab/fieldsensing/server/interfaces/database"
	"github.com/Pluslab/fieldsensing/server/usecase"
)

// IndividualController model
type IndividualController struct {
	Interactor usecase.IndividualInteractor
}

// NewIndividualController return the IndividualController
func NewIndividualController(SQLHandler database.SQLHandler) *IndividualController {
	return &IndividualController{
		Interactor: usecase.IndividualInteractor{
			IndividualRepository: &database.IndividualRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *IndividualController) Create(c Context) {
	f := domain.Individual{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *IndividualController) Index(c Context) {
	fields, err := controller.Interactor.Individuals()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *IndividualController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.IndividualByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
