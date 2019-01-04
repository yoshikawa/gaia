package usecase

import "github.com/Pluslab/gaia/server/domain"

// ObservationDatumInteractor model
type ObservationDatumInteractor struct {
	ObservationDatumRepository ObservationDatumRepository
}

// Add is to store observation data into database
func (interactor *ObservationDatumInteractor) Add(u domain.ObservationDatum) (observationDatum domain.ObservationDatum, err error) {
	identifier, err := interactor.ObservationDatumRepository.Store(u)
	if err != nil {
		return
	}
	observationDatum, err = interactor.ObservationDatumRepository.FindByID(identifier)
	return
}

// ObservationData is to find all observation data from database
func (interactor *ObservationDatumInteractor) ObservationData() (observationDatum domain.ObservationData, err error) {
	observationDatum, err = interactor.ObservationDatumRepository.FindAll()
	return
}

// ObservationDatumByID is to find by observation datum's id from database
func (interactor *ObservationDatumInteractor) ObservationDatumByID(identifier int64) (observationDatum domain.ObservationDatum, err error) {
	observationDatum, err = interactor.ObservationDatumRepository.FindByID(identifier)
	return
}
