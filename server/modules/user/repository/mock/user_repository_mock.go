package mock

import (
	organizationDomain "github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/Pluslab/gaia/server/modules/user/domain"
)

// UserRepositoryMock struct
type UserRepositoryMock struct {
	db map[uint]*domain.User
}

// NewUserRepositoryMock function
func NewUserRepositoryMock() *UserRepositoryMock {
	db := make(map[uint]*domain.User)
	db[1] = &domain.User{
		Name:           "yoshikawa",
		Email:          "yoshikawataiki@gmail.com",
		Password:       "hogehoge",
		Country:        "Japan",
		OrganizationID: 1,
		Organization: organizationDomain.Organization{
			Name: "pluslab",
		},
	}

	db[2] = &domain.User{
		Name:           "test",
		Email:          "test@gmail.com",
		Password:       "hogehuga",
		Country:        "Japan",
		OrganizationID: 2,
		Organization: organizationDomain.Organization{
			Name: "test2",
		},
	}
	return &UserRepositoryMock{db: db}
}

// Save function
func (userRepository *UserRepositoryMock) Save(user *domain.User) shared.Output {
	userRepository.db[user.ID] = user
	return shared.Output{Result: user}
}
