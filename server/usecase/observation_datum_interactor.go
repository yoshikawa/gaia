package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// ObservationDatumInteractor model
type ObservationDatumInteractor struct {
	ObservationDatumRepository ObservationDatumRepository
}

// Add is to store gateway device's data into database
func (interactor *ObservationDatumInteractor) Add(u domain.ObservationDatum) (observationDatum domain.ObservationDatum, err error) {
	identifier, err := interactor.ObservationDatumRepository.Store(u)
	if err != nil {
		return
	}
	observationDatum, err = interactor.ObservationDatumRepository.FindByID(identifier)
	return
}

// ObservationData is to find all gateway device's data from database
func (interactor *ObservationDatumInteractor) ObservationData() (observationDatum domain.ObservationData, err error) {
	observationDatum, err = interactor.ObservationDatumRepository.FindAll()
	return
}

// ObservationDatumByID is to find by gateway device's id from database
func (interactor *ObservationDatumInteractor) ObservationDatumByID(identifier int64) (observationDatum domain.ObservationDatum, err error) {
	observationDatum, err = interactor.ObservationDatumRepository.FindByID(identifier)
	return
}
