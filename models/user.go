package models

import "time"

// User Model
type User struct {
	Id             int64     `db:"id" json:"id"`
	OrganizationID int64     `db:"organization_id" json:"organization_id"`
	Name           string    `db:"name" json:"name"`
	Email          string    `db:"email" json:"email"`
	Password       string    `db:"password" json:"pasword"`
	Country        string    `db:"country" json:"country"`
	Administrator  int       `db:"administrator" json:"administrator"`
	createdAt      time.Time `db:"created_at" json:"created_at"`
	updatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
