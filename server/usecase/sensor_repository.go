package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// SensorRepository model
type SensorRepository interface {
	Store(domain.Sensor) (int64, error)
	FindByID(int64) (domain.Sensor, error)
	FindAll() (domain.Sensors, error)
}
