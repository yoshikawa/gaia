package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
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

// Create is to create the individual
func (controller *IndividualController) Create(c Context) {
	i := domain.Individual{}
	c.Bind(&i)
	individual, err := controller.Interactor.Add(i)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, individual)
}

// Index is to index the individuals
func (controller *IndividualController) Index(c Context) {
	individuals, err := controller.Interactor.Individuals()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, individuals)
}

// Show is to show the individual
func (controller *IndividualController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	individual, err := controller.Interactor.IndividualByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, individual)
}
