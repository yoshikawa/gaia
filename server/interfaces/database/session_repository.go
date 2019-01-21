package database

import (
	"github.com/Pluslab/gaia/server/domain"
)

// SessionRepository model
type SessionRepository struct {
	SQLHandler
}

// Login find the user by user email and password
func (repo *SessionRepository) Login(inputEmail string) (user domain.User, err error) {
	row, err := repo.Query("SELECT password FROM users WHERE email = ?", inputEmail)
	defer row.Close()
	if err != nil {
		return
	}
	var registeredPassword string
	row.Next()
	if err = row.Scan(&registeredPassword); err != nil {
		return
	}
	user.Email = inputEmail
	user.Password = registeredPassword
	return
}
