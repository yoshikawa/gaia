package usecase

import "github.com/Pluslab/gaia/server/domain"

// ObservablePropertyRepository model
type ObservablePropertyRepository interface {
	Store(domain.ObservableProperty) (int64, error)
	FindByID(int64) (domain.ObservableProperty, error)
	FindAll() (domain.ObservableProperties, error)
}
