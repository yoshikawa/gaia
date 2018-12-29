package usecase

import "github.com/Pluslab/gaia/server/domain"

// IndividualRepository model
type IndividualRepository interface {
	Store(domain.Individual) (int64, error)
	FindByID(int64) (domain.Individual, error)
	FindAll() (domain.Individuals, error)
}
