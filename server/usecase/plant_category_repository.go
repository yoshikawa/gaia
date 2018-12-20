package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// PlantCategoryRepository model
type PlantCategoryRepository interface {
	Store(domain.PlantCategory) (int64, error)
	FindByID(int64) (domain.PlantCategory, error)
	FindAll() (domain.PlantCategories, error)
}
