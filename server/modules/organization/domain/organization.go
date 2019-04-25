package domain

import (
	"errors"

	"github.com/Pluslab/gaia/server/modules/shared"
)

// Organization Model
type Organization struct {
	ID   uint   `gorm:"column:id; primary_key:yes"`
	Name string `gorm:"column:name; not null"`
	shared.BaseDomain
}

// Organizations type list of Organization
type Organizations []*Organization

// TableName function
func (organization Organization) TableName() string {
	return "organizations"
}

// Validate function
func (organization *Organization) Validate() error {
	if len(organization.Name) <= 0 {
		return errors.New("organization name is required")
	}
	return nil
}
