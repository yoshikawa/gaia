package model

// User Model
type User struct {
	ID   int64  `db:"id, primarykey"`
	Name string `json:"name"`
}
