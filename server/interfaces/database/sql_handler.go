package database

// SQLHandler Model
type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

// Result Model
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Row Model
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
