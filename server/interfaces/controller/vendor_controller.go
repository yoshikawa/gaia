package controller

import (
	"strconv"

	"github.com/Pluslab/gaia/server/domain"
	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

// VendorController model
type VendorController struct {
	Interactor usecase.VendorInteractor
}

// NewVendorController return the VendorController
func NewVendorController(SQLHandler database.SQLHandler) *VendorController {
	return &VendorController{
		Interactor: usecase.VendorInteractor{
			VendorRepository: &database.VendorRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the field
func (controller *VendorController) Create(c Context) {
	f := domain.Vendor{}
	c.Bind(&f)
	field, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, field)
}

// Index is to index the fields
func (controller *VendorController) Index(c Context) {
	fields, err := controller.Interactor.Vendors()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, fields)
}

// Show is to show the field
func (controller *VendorController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	field, err := controller.Interactor.VendorByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, field)
}
