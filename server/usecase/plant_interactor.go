package usecase

import "github.com/Pluslab/gaia/server/domain"

// PlantInteractor model
type PlantInteractor struct {
	PlantRepository PlantRepository
}

// Add is to store plant's data into database
func (interactor *PlantInteractor) Add(u domain.Plant) (plant domain.Plant, err error) {
	identifier, err := interactor.PlantRepository.Store(u)
	if err != nil {
		return
	}
	plant, err = interactor.PlantRepository.FindByID(identifier)
	return
}

// Plants is to find all plant's data from database
func (interactor *PlantInteractor) Plants() (plants domain.Plants, err error) {
	plants, err = interactor.PlantRepository.FindAll()
	return
}

// PlantByID is to find by plant's id from database
func (interactor *PlantInteractor) PlantByID(identifier int64) (plant domain.Plant, err error) {
	plant, err = interactor.PlantRepository.FindByID(identifier)
	return
}
