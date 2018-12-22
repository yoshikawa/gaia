package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// VendorRepository model
type VendorRepository interface {
	Store(domain.Vendor) (int64, error)
	FindByID(int64) (domain.Vendor, error)
	FindAll() (domain.Vendors, error)
}
