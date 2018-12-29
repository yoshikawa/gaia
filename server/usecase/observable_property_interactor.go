package usecase

import "github.com/Pluslab/gaia/server/domain"

// ObservablePropertyInteractor model
type ObservablePropertyInteractor struct {
	ObservablePropertyRepository ObservablePropertyRepository
}

// Add is to store observable property's data into database
func (interactor *ObservablePropertyInteractor) Add(u domain.ObservableProperty) (observableProperty domain.ObservableProperty, err error) {
	identifier, err := interactor.ObservablePropertyRepository.Store(u)
	if err != nil {
		return
	}
	observableProperty, err = interactor.ObservablePropertyRepository.FindByID(identifier)
	return
}

// ObservableProperties is to find all observable property's data from database
func (interactor *ObservablePropertyInteractor) ObservableProperties() (observableProperties domain.ObservableProperties, err error) {
	observableProperties, err = interactor.ObservablePropertyRepository.FindAll()
	return
}

// ObservablePropertyByID is to find by observable property's id from database
func (interactor *ObservablePropertyInteractor) ObservablePropertyByID(identifier int64) (observableProperty domain.ObservableProperty, err error) {
	observableProperty, err = interactor.ObservablePropertyRepository.FindByID(identifier)
	return
}
