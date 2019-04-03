package repository

import (
	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
)

// OrganizationRepository interface
type OrganizationRepository interface {
	Save(*domain.Organization) shared.Output
	Delete(*domain.Organization) shared.Output
	FindByID(uint) shared.Output
	FindAll(*shared.Parameters) shared.Output
	Count(*shared.Parameters) shared.Output
}
