package usecase

import "github.com/Pluslab/gaia/server/domain"

// SensorInteractor model
type SensorInteractor struct {
	SensorRepository SensorRepository
}

// Add is to store sensor's data into database
func (interactor *SensorInteractor) Add(u domain.Sensor) (sensor domain.Sensor, err error) {
	identifier, err := interactor.SensorRepository.Store(u)
	if err != nil {
		return
	}
	sensor, err = interactor.SensorRepository.FindByID(identifier)
	return
}

// Sensors is to find all sensor's data from database
func (interactor *SensorInteractor) Sensors() (sensors domain.Sensors, err error) {
	sensors, err = interactor.SensorRepository.FindAll()
	return
}

// SensorByID is to find by sensor's id from database
func (interactor *SensorInteractor) SensorByID(identifier int64) (sensor domain.Sensor, err error) {
	sensor, err = interactor.SensorRepository.FindByID(identifier)
	return
}
