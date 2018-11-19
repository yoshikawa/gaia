package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// UserInteractor model
type UserInteractor struct {
	UserRepository UserRepository
}

// Add ...
func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindByID(identifier)
	return
}

// Users ...
func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

// UserByID ...
func (interactor *UserInteractor) UserByID(identifier int64) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(identifier)
	return
}
