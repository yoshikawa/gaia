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
	row, err := repo.Query("SELECT id, name, password FROM users WHERE email = ? LIMIT 1", inputEmail)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var name string
	var password string
	row.Next()
	if err = row.Scan(&id, &name, &password); err != nil {
		return
	}
	user.ID = id
	user.Name = name
	user.Password = password
	return
}

// Logout is a function to wrap session repo
func (repo *SessionRepository) Logout() {
	return
}
