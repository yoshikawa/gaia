package domain

import (
	"errors"
	"time"

	"github.com/Pluslab/gaia/server/modules/organization/domain"
)

// User Model
type User struct {
	ID             uint `gorm:"column:id; primary_key:yes"`
	OrganizationID uint `gorm:"column:organization_id"`
	Organization   domain.Organization
	Name           string    `gorm:"column:name"`
	Email          string    `gorm:"column:email; unique_index"`
	Password       string    `gorm:"column:password; not null"`
	Country        string    `gorm:"column:country"`
	Administrator  bool      `gorm:"column:administrator"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

// Users type list of User
type Users []*User

// Validate function
func (user *User) Validate() error {
	if len(user.Name) <= 0 {
		return errors.New("user name is required")
	}

	if len(user.Email) <= 0 {
		return errors.New("user email is required")
	}

	if len(user.Password) <= 0 {
		return errors.New("user password is required")
	}

	return nil
}
