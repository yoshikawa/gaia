package usecase

import "github.com/Pluslab/fieldsensing/server/domain"

// PlantRepository model
type PlantRepository interface {
	Store(domain.Plant) (int64, error)
	FindByID(int64) (domain.Plant, error)
	FindAll() (domain.Plants, error)
}
