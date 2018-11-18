package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// OrganizationRepository interface model
// Input from interfaces/database/organization_repository
type OrganizationRepository interface {
	Store(domain.Organization) (int64, error)
	FindByID(int64) (domain.Organization, error)
	FindAll() (domain.Organizations, error)
}
