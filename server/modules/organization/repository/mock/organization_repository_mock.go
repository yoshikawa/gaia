package mock

import (
	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/jinzhu/gorm"
)

// OrganizationRepositoryMock struct
type OrganizationRepositoryMock struct {
	db map[uint]*domain.Organization
}

// NewOrganizationRepositoryMock function
func NewOrganizationRepositoryMock() *OrganizationRepositoryMock {
	db := make(map[uint]*domain.Organization)
	db[1] = &domain.Organization{
		Name: "pluslab",
	}

	db[2] = &domain.Organization{
		Name: "gaia",
	}
	return &OrganizationRepositoryMock{db: db}
}

// Save function
func (organizationRepositoryMock *OrganizationRepositoryMock) Save(organization *domain.Organization) shared.Output {
	organizationRepositoryMock.db[organization.ID] = organization
	return shared.Output{Result: organization}
}

// Delete function
func (organizationRepositoryMock *OrganizationRepositoryMock) Delete(organization *domain.Organization) shared.Output {
	organization, ok := organizationRepositoryMock.db[organization.ID]
	if !ok {
		return shared.Output{Err: gorm.ErrRecordNotFound}
	}

	delete(organizationRepositoryMock.db, organization.ID)

	return shared.Output{Result: organization}
}

// FindByID function
func (organizationRepositoryMock *OrganizationRepositoryMock) FindByID(id uint) shared.Output {
	organization, ok := organizationRepositoryMock.db[id]
	if !ok {
		return shared.Output{Err: gorm.ErrRecordNotFound}
	}

	return shared.Output{Result: organization}
}

// FindAll function
func (organizationRepositoryMock *OrganizationRepositoryMock) FindAll(params *shared.Parameters) shared.Output {
	var organizations domain.Organizations

	for _, organization := range organizations {
		organizations = append(organizations, organization)
	}

	return shared.Output{Result: organizations}
}

// Count function
func (organizationRepositoryMock *OrganizationRepositoryMock) Count(params *shared.Parameters) shared.Output {
	return shared.Output{Result: len(organizationRepositoryMock.db)}
}
