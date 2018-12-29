package usecase

import "github.com/Pluslab/gaia/server/domain"

// IndividualInteractor model
type IndividualInteractor struct {
	IndividualRepository IndividualRepository
}

// Add is to store gateway device's data into database
func (interactor *IndividualInteractor) Add(u domain.Individual) (individual domain.Individual, err error) {
	identifier, err := interactor.IndividualRepository.Store(u)
	if err != nil {
		return
	}
	individual, err = interactor.IndividualRepository.FindByID(identifier)
	return
}

// Individuals is to find all gateway device's data from database
func (interactor *IndividualInteractor) Individuals() (individuals domain.Individuals, err error) {
	individuals, err = interactor.IndividualRepository.FindAll()
	return
}

// IndividualByID is to find by gateway device's id from database
func (interactor *IndividualInteractor) IndividualByID(identifier int64) (individual domain.Individual, err error) {
	individual, err = interactor.IndividualRepository.FindByID(identifier)
	return
}
