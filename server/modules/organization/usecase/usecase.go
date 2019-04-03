package usecase

import (
	"github.com/Pluslab/gaia/server/modules/organization/domain"
	"github.com/Pluslab/gaia/server/modules/shared"
)

// ProductUsecase interface
type ProductUsecase interface {
	CreateProduct(*domain.Organization) shared.Output
	GetProduct(uint) shared.Output
	GetAllProduct(*shared.Parameters) shared.Output
	GetTotalProduct(*shared.Parameters) shared.Output
	RemoveProduct(uint) shared.Output
}
