package usecase

import (
	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
)

// OrganizationUsecase interface
type OrganizationUsecase interface {
	CreateOrganization(*domain.Organization) shared.Output
	GetOrganization(uint) shared.Output
	GetAllOrganization(*shared.Parameters) shared.Output
	GetTotalOrganization(*shared.Parameters) shared.Output
	RemoveOrganization(uint) shared.Output
}
