package domain

import (
	"errors"

	"github.com/paulmach/orb/geo"

	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/paulmach/orb"
)

// Field Model
type Field struct {
	ID     int64       `gorm:"column:id; primary_key:yes"`
	Groups string      `gorm:"column:groups"`
	Name   string      `gorm:"column:name"`
	Area   orb.Polygon `gorm:"column:area"`
	shared.BaseDomain
	Deleted bool `gorm:"column:deleted"`
}

// Fields type list of Field
type Fields []*Field

// TableName function
func (field Field) TableName() string {
	return "fields"
}

// Validate function
func (field *Field) Validate() error {
	if len(field.Name) <= 0 {
		return errors.New("field name is required")
	}

	if geo.Area(field.Area) <= 0 {
		return errors.New("field area is required")
	}

	return nil
}
