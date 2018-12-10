package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// FieldRepository model
type FieldRepository interface {
	Store(domain.Field) (int64, error)
	FindByID(int64) (domain.Field, error)
	FindAll() (domain.Fields, error)
}
