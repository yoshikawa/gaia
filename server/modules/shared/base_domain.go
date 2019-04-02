package shared

import "time"

// BaseDomain structure
type BaseDomain struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
