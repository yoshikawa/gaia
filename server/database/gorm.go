package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// GetGormConn function
func GetGormConn(host, user, dbName, password string, port int) (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		host, port, user, dbName, password,
	))
}
