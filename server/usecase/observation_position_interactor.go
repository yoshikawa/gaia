package usecase

import "github.com/Pluslab/gaia/server/domain"

// ObservationPositionInteractor model
type ObservationPositionInteractor struct {
	ObservationPositionRepository ObservationPositionRepository
}

// Add is to store gateway device's data into database
func (interactor *ObservationPositionInteractor) Add(u domain.ObservationPosition) (observationPosition domain.ObservationPosition, err error) {
	identifier, err := interactor.ObservationPositionRepository.Store(u)
	if err != nil {
		return
	}
	observationPosition, err = interactor.ObservationPositionRepository.FindByID(identifier)
	return
}

// ObservationPositions is to find all gateway device's data from database
func (interactor *ObservationPositionInteractor) ObservationPositions() (observationPositions domain.ObservationPositions, err error) {
	observationPositions, err = interactor.ObservationPositionRepository.FindAll()
	return
}

// ObservationPositionByID is to find by gateway device's id from database
func (interactor *ObservationPositionInteractor) ObservationPositionByID(identifier int64) (observationPosition domain.ObservationPosition, err error) {
	observationPosition, err = interactor.ObservationPositionRepository.FindByID(identifier)
	return
}
