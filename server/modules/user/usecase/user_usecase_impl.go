package usecase

import (
	organizationRepository "github.com/Pluslab/gaia/server/modules/organization/repository"
	"github.com/Pluslab/gaia/server/modules/user/repository"
)

// UserUsecaseImpl struct
type UserUsecaseImpl struct {
	userRepositoryRead     repository.UserRepository
	userRepositoryWrite    repository.UserRepository
	organizationRepository organizationRepository.OrganizationRepository
}
