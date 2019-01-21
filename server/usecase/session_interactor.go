package usecase

import "github.com/Pluslab/gaia/server/domain"

// SessionInteractor model
type SessionInteractor struct {
	SessionRepository SessionRepository
}

// Login is to login by user's email from database
func (interactor *SessionInteractor) Login(email string) (user domain.User, err error) {
	user, err = interactor.SessionRepository.Login(email)
	return
}

// Logout is a function for logging out
func (interactor *SessionInteractor) Logout() {
	interactor.SessionRepository.Logout()
}
