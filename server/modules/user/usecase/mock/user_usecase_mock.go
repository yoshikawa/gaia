package mock

import (
	organizationDomain "github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/Pluslab/gaia/server/modules/user/domain"
)

// UserUsecaseMock struct
type UserUsecaseMock struct {
}

// NewUserUsecaseMock function
func NewUserUsecaseMock() *UserUsecaseMock {
	return &UserUsecaseMock{}
}

// CreateUser function
func (userUsecase *UserUsecaseMock) CreateUser(user *domain.User) shared.Output {
	userSaved := &domain.User{
		Name:           "yoshikawa",
		Email:          "yoshikawataiki@gmail.com",
		Password:       "hogehoge",
		Country:        "Japan",
		OrganizationID: 1,
		Organization: organizationDomain.Organization{
			Name: "pluslab",
		},
	}
	return shared.Output{Result: userSaved}
}

// RemoveUser function
func (userUsecase *UserUsecaseMock) RemoveUser(id uint) shared.Output {
	user := &domain.User{
		Name:           "yoshikawa",
		Email:          "yoshikawataiki@gmail.com",
		Password:       "hogehoge",
		Country:        "Japan",
		OrganizationID: 1,
		Organization: organizationDomain.Organization{
			Name: "pluslab",
		},
	}
	return shared.Output{Result: user}
}
