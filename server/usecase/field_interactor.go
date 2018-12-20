package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// FieldInteractor model
type FieldInteractor struct {
	FieldRepository FieldRepository
}

// Add is to add field data
func (interactor *FieldInteractor) Add(f domain.Field) (field domain.Field, err error) {
	identifier, err := interactor.FieldRepository.Store(f)
	if err != nil {
		return
	}
	field, err = interactor.FieldRepository.FindByID(identifier)
	return
}

// Fields is to find all fields
func (interactor *FieldInteractor) Fields() (fields domain.Fields, err error) {
	fields, err = interactor.FieldRepository.FindAll()
	return
}

// FieldByID is to find field by id
func (interactor *FieldInteractor) FieldByID(identifier int64) (field domain.Field, err error) {
	field, err = interactor.FieldRepository.FindByID(identifier)
	return
}
