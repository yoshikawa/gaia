package usecase

import (
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/Pluslab/gaia/server/modules/user/domain"
)

// UserUsecase interface
type UserUsecase interface {
	CreateUser(*domain.User) shared.Output
	GetUser(uint) shared.Output
	GetAllUser(*shared.Parameters) shared.Output
	GetTotalUser(*shared.Parameters) shared.Output
	RemoveUser(uint) shared.Output
}
