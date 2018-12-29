package usecase

import "github.com/Pluslab/gaia/server/domain"

// GatewayDeviceInteractor model
type GatewayDeviceInteractor struct {
	GatewayDeviceRepository GatewayDeviceRepository
}

// Add is to store gateway device's data into database
func (interactor *GatewayDeviceInteractor) Add(u domain.GatewayDevice) (gatewayDevice domain.GatewayDevice, err error) {
	identifier, err := interactor.GatewayDeviceRepository.Store(u)
	if err != nil {
		return
	}
	gatewayDevice, err = interactor.GatewayDeviceRepository.FindByID(identifier)
	return
}

// GatewayDevices is to find all gateway device's data from database
func (interactor *GatewayDeviceInteractor) GatewayDevices() (gatewayDevice domain.GatewayDevices, err error) {
	gatewayDevice, err = interactor.GatewayDeviceRepository.FindAll()
	return
}

// GatewayDeviceByID is to find by gateway device's id from database
func (interactor *GatewayDeviceInteractor) GatewayDeviceByID(identifier int64) (gatewayDevice domain.GatewayDevice, err error) {
	gatewayDevice, err = interactor.GatewayDeviceRepository.FindByID(identifier)
	return
}
