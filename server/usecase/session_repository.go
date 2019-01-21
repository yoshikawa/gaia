package usecase

import "github.com/Pluslab/gaia/server/domain"

// SessionRepository model
type SessionRepository interface {
	Login(string) (domain.User, error)
}
