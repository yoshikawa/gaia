package usecase

import "github.com/Pluslab/gaia/server/domain"

// GatewayDeviceRepository model
type GatewayDeviceRepository interface {
	Store(domain.GatewayDevice) (int64, error)
	FindByID(int64) (domain.GatewayDevice, error)
	FindAll() (domain.GatewayDevices, error)
}
