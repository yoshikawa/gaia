package infrastructure

import (
	"database/sql"

	"github.com/Pluslab/fieldsensing/server/interfaces/database"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler model
type SQLHandler struct {
	Conn *sql.DB
}

// SQLResult model
type SQLResult struct {
	Result sql.Result
}

// SQLRow model
type SQLRow struct {
	Rows *sql.Rows
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

// Execute is to execute the commands fot the database
func (handler *SQLHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

// Query is to issue a query
func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

// LastInsertId is to insert the latest id
func (r SQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected is to affect the rows
func (r SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// Scan is to scan dest
func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next is to check if there is next
func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

// Close is to close sql
func (r SQLRow) Close() error {
	return r.Rows.Close()
}
