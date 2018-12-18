package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// GatewayDeviceRepository model
type GatewayDeviceRepository interface {
	Store(domain.GatewayDevice) (int64, error)
	FindByID(int64) (domain.GatewayDevice, error)
	FindAll() (domain.GatewayDevices, error)
}
