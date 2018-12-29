package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// FieldController model
type FieldController struct {
	Interactor usecase.FieldInteractor
}

// NewFieldController return the FieldController
func NewFieldController(SQLHandler database.SQLHandler) *FieldController {
	return &FieldController{
		Interactor: usecase.FieldInteractor{
			FieldRepository: &database.FieldRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *FieldController) Create(c Context) {
	f := domain.Field{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *FieldController) Index(c Context) {
	fields, err := controller.Interactor.Fields()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *FieldController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.FieldByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
