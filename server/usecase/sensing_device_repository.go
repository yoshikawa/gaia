package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// SensingDeviceRepository model
type SensingDeviceRepository interface {
	Store(domain.SensingDevice) (int64, error)
	FindByID(int64) (domain.SensingDevice, error)
	FindAll() (domain.SensingDevices, error)
}
