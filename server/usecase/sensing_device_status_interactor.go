package usecase

import "github.com/Pluslab/gaia/server/domain"

// SensingDeviceStatusInteractor model
type SensingDeviceStatusInteractor struct {
	SensingDeviceStatusRepository SensingDeviceStatusRepository
}

// Add is to store sensing device status's data into database
func (interactor *SensingDeviceStatusInteractor) Add(u domain.SensingDeviceStatus) (sensingDeviceStatus domain.SensingDeviceStatus, err error) {
	identifier, err := interactor.SensingDeviceStatusRepository.Store(u)
	if err != nil {
		return
	}
	sensingDeviceStatus, err = interactor.SensingDeviceStatusRepository.FindByID(identifier)
	return
}

// SensingDeviceStatuses is to find all sensing device statu's data from database
func (interactor *SensingDeviceStatusInteractor) SensingDeviceStatuses() (sensingDeviceStatuses domain.SensingDeviceStatuses, err error) {
	sensingDeviceStatuses, err = interactor.SensingDeviceStatusRepository.FindAll()
	return
}

// SensingDeviceStatusByID is to find by sensing status's id from database
func (interactor *SensingDeviceStatusInteractor) SensingDeviceStatusByID(identifier int64) (sensingDeviceStatus domain.SensingDeviceStatus, err error) {
	sensingDeviceStatus, err = interactor.SensingDeviceStatusRepository.FindByID(identifier)
	return
}
