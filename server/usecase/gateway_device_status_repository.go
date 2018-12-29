package usecase

import "github.com/Pluslab/gaia/server/domain"

// GatewayDeviceStatusRepository model
type GatewayDeviceStatusRepository interface {
	Store(domain.GatewayDeviceStatus) (int64, error)
	FindByID(int64) (domain.GatewayDeviceStatus, error)
	FindAll() (domain.GatewayDeviceStatuses, error)
}
