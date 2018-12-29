package usecase

import "github.com/Pluslab/gaia/server/domain"

// SensingDeviceInteractor model
type SensingDeviceInteractor struct {
	SensingDeviceRepository SensingDeviceRepository
}

// Add is to store sensing device's data into database
func (interactor *SensingDeviceInteractor) Add(u domain.SensingDevice) (sensingDevice domain.SensingDevice, err error) {
	identifier, err := interactor.SensingDeviceRepository.Store(u)
	if err != nil {
		return
	}
	sensingDevice, err = interactor.SensingDeviceRepository.FindByID(identifier)
	return
}

// SensingDevices is to find all sensing device's data from database
func (interactor *SensingDeviceInteractor) SensingDevices() (sensingDevices domain.SensingDevices, err error) {
	sensingDevices, err = interactor.SensingDeviceRepository.FindAll()
	return
}

// SensingDeviceByID is to find by sensing device's id from database
func (interactor *SensingDeviceInteractor) SensingDeviceByID(identifier int64) (sensingDevice domain.SensingDevice, err error) {
	sensingDevice, err = interactor.SensingDeviceRepository.FindByID(identifier)
	return
}
