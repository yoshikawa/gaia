package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// ObservationPositionRepository model
type ObservationPositionRepository interface {
	Store(domain.ObservationPosition) (int64, error)
	FindByID(int64) (domain.ObservationPosition, error)
	FindAll() (domain.ObservationPositions, error)
}
