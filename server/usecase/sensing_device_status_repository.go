package usecase

import "github.com/Pluslab/gaia/server/domain"

// SensingDeviceStatusRepository model
type SensingDeviceStatusRepository interface {
	Store(domain.SensingDeviceStatus) (int64, error)
	FindByID(int64) (domain.SensingDeviceStatus, error)
	FindAll() (domain.SensingDeviceStatuses, error)
}
