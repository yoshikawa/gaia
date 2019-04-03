package repository

import (
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/Pluslab/gaia/server/modules/user/domain"
)

// UserRepository interface
type UserRepository interface {
	Save(*domain.User) shared.Output
	Delete(*domain.User) shared.Output
	FindByID(string) shared.Output
	FindAll(*shared.Parameters) shared.Output
	Count(*shared.Parameters) shared.Output
}
