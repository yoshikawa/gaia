package usecase

import (
	"fmt"

	organizationRepository "github.com/Pluslab/gaia/server/modules/organization/repository"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/Pluslab/gaia/server/modules/user/domain"
	"github.com/Pluslab/gaia/server/modules/user/repository"
	"github.com/jinzhu/gorm"
)

// UserUsecaseImpl struct
type UserUsecaseImpl struct {
	userRepositoryRead     repository.UserRepository
	userRepositoryWrite    repository.UserRepository
	organizationRepository organizationRepository.OrganizationRepository
}

// NewUserUsecaseImpl function
func NewUserUsecaseImpl(userRepositoryRead, userRepositoryWrite repository.UserRepository,
	organizationRepository organizationRepository.OrganizationRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		userRepositoryRead:     userRepositoryRead,
		userRepositoryWrite:    userRepositoryWrite,
		organizationRepository: organizationRepository,
	}
}

// CreateUser function
func (u *UserUsecaseImpl) CreateUser(user *domain.User) shared.Output {
	userOutput := u.userRepositoryRead.FindByID(user.ID)
	if userOutput.Err != nil && userOutput.Err != gorm.ErrRecordNotFound {
		return shared.Output{Err: userOutput.Err}
	}

	if userOutput.Result != nil {
		userExist := userOutput.Result.(*domain.User)

		if userExist != nil {
			return shared.Output{Err: fmt.Errorf("user with id: %d already exist", user.ID)}
		}
	}

	err := user.Validate()

	if err != nil {
		return shared.Output{Err: err}
	}

	userSaveOutput := u.userRepositoryWrite.Save(user)
	if userSaveOutput.Err != nil {
		return shared.Output{Err: userSaveOutput.Err}
	}

	userSaved := userSaveOutput.Result.(*domain.User)

	return shared.Output{Result: userSaved}
}

// RemoveUser function
func (u *UserUsecaseImpl) RemoveUser(id uint) shared.Output {
	userResult := u.userRepositoryRead.FindByID(id)
	if userResult.Err != nil {
		return shared.Output{Err: userResult.Err}
	}

	user := userResult.Result.(*domain.User)

	userDeleteOutput := u.userRepositoryWrite.Delete(user)
	if userDeleteOutput.Err != nil {
		return shared.Output{Err: userDeleteOutput.Err}
	}

	userDeleted := userDeleteOutput.Result.(*domain.User)
	return shared.Output{Result: userDeleted}
}
