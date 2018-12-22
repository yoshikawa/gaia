package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// ObservationDatumRepository model
type ObservationDatumRepository interface {
	Store(domain.ObservationDatum) (int64, error)
	FindByID(int64) (domain.ObservationDatum, error)
	FindAll() (domain.ObservationData, error)
}
