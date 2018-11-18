package usecase

import (
	"github.com/Pluslab/fieldsensing/server/domain"
)

// OrganizationInteractor model
type OrganizationInteractor struct {
	OrganizationRepository OrganizationRepository
}

// Add is to add organization
// this is Gateway. input port.
func (interactor *OrganizationInteractor) Add(organization domain.Organization) (err error) {
	_, err = interactor.OrganizationRepository.Store(organization)
	return
}

// Organizations is to find all organizations
func (interactor *OrganizationInteractor) Organizations() (organizations domain.Organizations, err error) {
	organizations, err = interactor.OrganizationRepository.FindAll()
	return
}

// OrganizationByID is to find the organization by ID
func (interactor *OrganizationInteractor) OrganizationByID(identifier int64) (organization domain.Organization, err error) {
	organization, err = interactor.OrganizationRepository.FindByID(identifier)
	return
}
