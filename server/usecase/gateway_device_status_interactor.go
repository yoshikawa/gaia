package usecase

import "github.com/Pluslab/gaia/server/domain"

// GatewayDeviceStatusInteractor model
type GatewayDeviceStatusInteractor struct {
	GatewayDeviceStatusRepository GatewayDeviceStatusRepository
}

// Add is to store gateway device status's data into database
func (interactor *GatewayDeviceStatusInteractor) Add(u domain.GatewayDeviceStatus) (gatewayDeviceStatus domain.GatewayDeviceStatus, err error) {
	identifier, err := interactor.GatewayDeviceStatusRepository.Store(u)
	if err != nil {
		return
	}
	gatewayDeviceStatus, err = interactor.GatewayDeviceStatusRepository.FindByID(identifier)
	return
}

// GatewayDeviceStatuses is to find all gateway device status's data from database
func (interactor *GatewayDeviceStatusInteractor) GatewayDeviceStatuses() (gatewayDeviceStatuses domain.GatewayDeviceStatuses, err error) {
	gatewayDeviceStatuses, err = interactor.GatewayDeviceStatusRepository.FindAll()
	return
}

// GatewayDeviceStatusByID is to find by gateway device status's id from database
func (interactor *GatewayDeviceStatusInteractor) GatewayDeviceStatusByID(identifier int64) (gatewayDeviceStatus domain.GatewayDeviceStatus, err error) {
	gatewayDeviceStatus, err = interactor.GatewayDeviceStatusRepository.FindByID(identifier)
	return
}
