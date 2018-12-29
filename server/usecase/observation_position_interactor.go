package usecase

import "github.com/Pluslab/gaia/server/domain"

// ObservationPositionInteractor model
type ObservationPositionInteractor struct {
	ObservationPositionRepository ObservationPositionRepository
}

// Add is to store observation position's data into database
func (interactor *ObservationPositionInteractor) Add(u domain.ObservationPosition) (observationPosition domain.ObservationPosition, err error) {
	identifier, err := interactor.ObservationPositionRepository.Store(u)
	if err != nil {
		return
	}
	observationPosition, err = interactor.ObservationPositionRepository.FindByID(identifier)
	return
}

// ObservationPositions is to find all observation position's data from database
func (interactor *ObservationPositionInteractor) ObservationPositions() (observationPositions domain.ObservationPositions, err error) {
	observationPositions, err = interactor.ObservationPositionRepository.FindAll()
	return
}

// ObservationPositionByID is to find by observation position's id from database
func (interactor *ObservationPositionInteractor) ObservationPositionByID(identifier int64) (observationPosition domain.ObservationPosition, err error) {
	observationPosition, err = interactor.ObservationPositionRepository.FindByID(identifier)
	return
}
