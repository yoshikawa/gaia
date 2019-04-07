package mock

import (
	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
)

// OrganizationUsecaseMock struct
type OrganizationUsecaseMock struct {
}

// NewOrganizationUsecaseMock function
// return OrganizationUsecaseMock struct
func NewOrganizationUsecaseMock() *OrganizationUsecaseMock {
	return &OrganizationUsecaseMock{}
}

// CreateOrganization function
func (u *OrganizationUsecaseMock) CreateOrganization(organization *domain.Organization) shared.Output {
	organizationSaved := &domain.Organization{
		Name: "pluslab",
	}
	return shared.Output{Result: organizationSaved}
}

// RemoveOrganization function
func (u *OrganizationUsecaseMock) RemoveOrganization(id uint) shared.Output {
	organizationSaved := &domain.Organization{
		Name: "pluslab",
	}
	return shared.Output{Result: organizationSaved}
}

// GetOrganization function
func (u *OrganizationUsecaseMock) GetOrganization(id uint) shared.Output {
	organization := &domain.Organization{
		Name: "pluslab",
	}

	return shared.Output{Result: organization}
}

// GetAllOrganization function
func (u *OrganizationUsecaseMock) GetAllOrganization(params *shared.Parameters) shared.Output {

	organizations := domain.Organizations{
		&domain.Organization{
			Name: "pluslab",
		},
		&domain.Organization{
			Name: "pluslab2",
		},
	}
	return shared.Output{Result: organizations}
}

// GetTotalOrganization function
func (u *OrganizationUsecaseMock) GetTotalOrganization(params *shared.Parameters) shared.Output {
	return shared.Output{Result: 2}
}
