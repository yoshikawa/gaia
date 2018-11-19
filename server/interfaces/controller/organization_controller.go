package controller

import (
	"strconv"

	"github.com/Pluslab/fieldsensing/server/domain"
	"github.com/Pluslab/fieldsensing/server/interfaces/database"
	"github.com/Pluslab/fieldsensing/server/usecase"
)

// OrganizationController model
type OrganizationController struct {
	Interactor usecase.OrganizationInteractor
}

// NewOrganizationController return the OrganizationController
func NewOrganizationController(SQLHandler database.SQLHandler) *OrganizationController {
	return &OrganizationController{
		Interactor: usecase.OrganizationInteractor{
			OrganizationRepository: &database.OrganizationRepository{
				SQLHandler: SQLHandler,
			},
		},
	}
}

// Create is to create the organization
func (controller *OrganizationController) Create(c Context) {
	o := domain.Organization{}
	c.Bind(&o)
	organization, err := controller.Interactor.Add(o)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, organization)
}

// Index is to index the orgnizations
func (controller *OrganizationController) Index(c Context) {
	organizations, err := controller.Interactor.Organizations()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, organizations)
}

// Show is to show the organization
func (controller *OrganizationController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	id64 := int64(id)
	organization, err := controller.Interactor.OrganizationByID(id64)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, organization)
}
