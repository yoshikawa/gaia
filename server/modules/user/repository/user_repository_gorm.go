package repository

import (
	"github.com/Pluslab/gaia/server/modules/shared"
	"github.com/Pluslab/gaia/server/modules/user/domain"
	"github.com/jinzhu/gorm"
)

// UserRepositoryGorm struct
type UserRepositoryGorm struct {
	db *gorm.DB
}

// NewUserRepositoryGorm function
func NewUserRepositoryGorm(db *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{db: db}
}

// Save function
func (r *UserRepositoryGorm) Save(user *domain.User) shared.Output {
	err := r.db.Save(user).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: user}
}

// Delete function
func (r *UserRepositoryGorm) Delete(user *domain.User) shared.Output {
	err := r.db.Delete(user).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: user}
}
