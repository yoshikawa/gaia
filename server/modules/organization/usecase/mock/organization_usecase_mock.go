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
