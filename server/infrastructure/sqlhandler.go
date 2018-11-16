package infrastructure

import (
	"database/sql"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler model
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler handling SQL
func NewSQLHandler() *SQLHandler {
	conn, err := sql.Open("mysql", "root:password@tcp(db:3306)/fieldsensing")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
