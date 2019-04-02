package database

import (
	"fmt"

	"github.com/Pluslab/gaia/server/config"

	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// GetGormConn function
func GetGormConn() (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s",
		config.DBUser, config.DBPassword, config.DBProtocol, config.DBHost, config.DBPort, config.DBName, config.DBCharset, config.DBParseTime))
}
