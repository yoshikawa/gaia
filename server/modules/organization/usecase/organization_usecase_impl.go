package usecase

import (
	"fmt"

	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/organization/repository"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/jinzhu/gorm"
)

// OrganizationUsecaseImpl struct
type OrganizationUsecaseImpl struct {
	organizationRepositoryRead  repository.OrganizationRepository
	organizationRepositoryWrite repository.OrganizationRepository
}

// NewOrganizationUsecaseImpl function
func NewOrganizationUsecaseImpl(OrganizationRepositoryRead, OrganizationRepositoryWrite repository.OrganizationRepository) *OrganizationUsecaseImpl {
	return &OrganizationUsecaseImpl{
		organizationRepositoryRead:  OrganizationRepositoryRead,
		organizationRepositoryWrite: OrganizationRepositoryWrite,
	}
}

// CreateOrganization function
func (u *OrganizationUsecaseImpl) CreateOrganization(organization *domain.Organization) shared.Output {
	organizationOutput := u.organizationRepositoryRead.FindByID(organization.ID)
	if organizationOutput.Err != nil && organizationOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output{Err: organizationOutput.Err}
	}

	if organizationOutput.Result != nil {
		organizationExist := organizationOutput.Result.(*domain.Organization)

		if organizationExist != nil {
			return shared.Output{Err: fmt.Errorf("organization with id: %d already exist", organization.ID)}
		}
	}

	err := organization.Validate()

	if err != nil {
		return shared.Output{Err: err}
	}

	organizationSaveOutput := u.organizationRepositoryWrite.Save(organization)
	if organizationSaveOutput.Err != nil {
		return shared.Output{Err: organizationSaveOutput.Err}
	}

	organizationSaved := organizationSaveOutput.Result.(*domain.Organization)

	return shared.Output{Result: organizationSaved}
}
