package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// UserRepository model
type UserRepository interface {
	Store(domain.User) (int64, error)
	FindByID(int64) (domain.User, error)
	FindAll() (domain.Users, error)
}
