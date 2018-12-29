package usecase

import "github.com/Pluslab/gaia/server/domain"

// PlantCategoryInteractor model
type PlantCategoryInteractor struct {
	PlantCategoryRepository PlantCategoryRepository
}

// Add is to store gateway device's data into database
func (interactor *PlantCategoryInteractor) Add(u domain.PlantCategory) (plantCategory domain.PlantCategory, err error) {
	identifier, err := interactor.PlantCategoryRepository.Store(u)
	if err != nil {
		return
	}
	plantCategory, err = interactor.PlantCategoryRepository.FindByID(identifier)
	return
}

// PlantCategories is to find all gateway device's data from database
func (interactor *PlantCategoryInteractor) PlantCategories() (plantCategories domain.PlantCategories, err error) {
	plantCategories, err = interactor.PlantCategoryRepository.FindAll()
	return
}

// PlantCategoryByID is to find by gateway device's id from database
func (interactor *PlantCategoryInteractor) PlantCategoryByID(identifier int64) (plantCategory domain.PlantCategory, err error) {
	plantCategory, err = interactor.PlantCategoryRepository.FindByID(identifier)
	return
}
