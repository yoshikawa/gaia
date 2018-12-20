package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// VendorInteractor model
type VendorInteractor struct {
	VendorRepository VendorRepository
}

// Add is to store gateway device's data into database
func (interactor *VendorInteractor) Add(u domain.Vendor) (vendor domain.Vendor, err error) {
	identifier, err := interactor.VendorRepository.Store(u)
	if err != nil {
		return
	}
	vendor, err = interactor.VendorRepository.FindByID(identifier)
	return
}

// Vendors is to find all gateway device's data from database
func (interactor *VendorInteractor) Vendors() (vendors domain.Vendors, err error) {
	vendors, err = interactor.VendorRepository.FindAll()
	return
}

// VendorByID is to find by gateway device's id from database
func (interactor *VendorInteractor) VendorByID(identifier int64) (vendor domain.Vendor, err error) {
	vendor, err = interactor.VendorRepository.FindByID(identifier)
	return
}
