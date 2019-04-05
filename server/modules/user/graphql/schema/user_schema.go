package schema

import (
	"time"

	"github.com/Pluslab/gaia/server/modules/user/domain"
	"github.com/graph-gophers/graphql-go"
)

// UserSchemaInput strcut
type UserSchemaInput struct {
	OrganizationID uint
	Name           string
	Email          string
	Password       string
	Country        string
	Administrator  bool
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

// UserSchema resolver
type UserSchema struct {
	User *domain.User
}

// OrganizationID function
func (user *UserSchema) OrganizationID() graphql.ID {
	return graphql.ID(user.User.OrganizationID)
}

// Name function
func (user *UserSchema) Name() string {
	return user.User.Name
}

// Email function
func (user *UserSchema) Email() string {
	return user.User.Email
}

// Password funciton
func (user *UserSchema) Password() string {
	return user.User.Password
}

// Country funciton
func (user *UserSchema) Country() string {
	return user.User.Country
}

// Administrator function
func (user *UserSchema) Administrator() bool {
	return user.User.Administrator
}

// CreatedAt function
func (user *UserSchema) CreatedAt() *graphql.Time {
	return &graphql.Time{Time: user.User.CreatedAt}
}

// UpdatedAt function
func (user *UserSchema) UpdatedAt() *graphql.Time {
	return &graphql.Time{Time: user.User.UpdatedAt}
}
