package mock

import "github.com/Pluslab/gaia/server/modules/organization/domain"

// OrganizationRepositoryMock struct
type OrganizationRepositoryMock struct {
	db map[string]*domain.Organization
}

// NewOrganizationRepositoryMock function
func NewOrganizationRepositoryMock() *OrganizationRepositoryMock {
	db := make(map[string]*domain.Organization)
	db["1"] = &domain.Organization{
		Name: "pluslab",
	}

	db["2"] = &domain.Organization{
		Name: "gaia",
	}
	return &OrganizationRepositoryMock{db: db}
}
