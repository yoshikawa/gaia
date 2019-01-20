package database

import (
	"time"

	"github.com/Pluslab/gaia/server/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository model
type UserRepository struct {
	SQLHandler
}

// Store insert values into user table
func (repo *UserRepository) Store(u domain.User) (id int64, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(password)
	result, err := repo.Execute(
		"INSERT INTO users (organization_id, name, email, password, country, administrator) VALUES (?, ?, ?, ?, ?, ?)", u.OrganizationID, u.Name, u.Email, u.Password, u.Country, u.Administrator,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int64(id64)
	return
}

// FindByID find the user by id
func (repo *UserRepository) FindByID(identifier int64) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, organization_id, name, email, country, administrator, created_at, updated_at FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var organizationID int64
	var name string
	var email string
	var country string
	var administrator bool
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &organizationID, &name, &email, &country, &administrator, &createdAt, &updatedAt); err != nil {
		return
	}
	user.ID = id
	user.OrganizationID = organizationID
	user.Name = name
	user.Email = email
	user.Country = country
	user.Administrator = administrator
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return
}

// FindAll find all users
func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, organization_id, name, email, country, administrator, created_at, updated_at FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var organizationID int64
		var name string
		var email string
		var country string
		var administrator bool
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &organizationID, &name, &email, &country, &administrator, &createdAt, &updatedAt); err != nil {
			continue
		}
		user := domain.User{
			ID:             id,
			OrganizationID: organizationID,
			Name:           name,
			Email:          email,
			Country:        country,
			Administrator:  administrator,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		}
		users = append(users, user)
	}
	return
}

// Login find the user by user email
func (repo *UserRepository) Login(inputEmail string) (user domain.User, err error) {
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
