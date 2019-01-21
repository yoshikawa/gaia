package usecase

import (
	"github.com/Pluslab/gaia/server/domain"
)

// UserInteractor model
type UserInteractor struct {
	UserRepository UserRepository
}

// Add is to store user's data into database
func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindByID(identifier)
	return
}

// UserByID is to find by user's id from database
func (interactor *UserInteractor) UserByID(identifier int64) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(identifier)
	return
}

// Users is to find all user's data from database
func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}
