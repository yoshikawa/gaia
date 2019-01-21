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

// Create is to create the vendor
func (controller *VendorController) Create(c Context) {
	v := domain.Vendor{}
	c.Bind(&v)
	vendor, err := controller.Interactor.Add(v)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, vendor)
}

// Index is to index the vendors
func (controller *VendorController) Index(c Context) {
	vendors, err := controller.Interactor.Vendors()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, vendors)
}

// Show is to show the vendor
func (controller *VendorController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	vendor, err := controller.Interactor.VendorByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, vendor)
}
