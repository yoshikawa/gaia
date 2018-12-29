package usecase

import "github.com/Pluslab/gaia/server/domain"

// ObservationDatumRepository model
type ObservationDatumRepository interface {
	Store(domain.ObservationDatum) (int64, error)
	FindByID(int64) (domain.ObservationDatum, error)
	FindAll() (domain.ObservationData, error)
}
