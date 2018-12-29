package infrastructure

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/interfaces/database"
	_ "github.com/go-sql-driver/mysql"
)

func TestNewSQLHandler(t *testing.T) {
	tests := []struct {
		name string
		want *SQLHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSQLHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSQLHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLHandler_Execute(t *testing.T) {
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		statement string
		args      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    database.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &SQLHandler{
				Conn: tt.fields.Conn,
			}
			got, err := handler.Execute(tt.args.statement, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLHandler.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLHandler.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLHandler_Query(t *testing.T) {
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		statement string
		args      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    database.Row
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &SQLHandler{
				Conn: tt.fields.Conn,
			}
			got, err := handler.Query(tt.args.statement, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLHandler.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLHandler.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLResult_LastInsertId(t *testing.T) {
	type fields struct {
		Result sql.Result
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SQLResult{
				Result: tt.fields.Result,
			}
			got, err := r.LastInsertId()
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLResult.LastInsertId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SQLResult.LastInsertId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLResult_RowsAffected(t *testing.T) {
	type fields struct {
		Result sql.Result
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SQLResult{
				Result: tt.fields.Result,
			}
			got, err := r.RowsAffected()
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLResult.RowsAffected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SQLResult.RowsAffected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLRow_Scan(t *testing.T) {
	type fields struct {
		Rows *sql.Rows
	}
	type args struct {
		dest []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SQLRow{
				Rows: tt.fields.Rows,
			}
			if err := r.Scan(tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("SQLRow.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLRow_Next(t *testing.T) {
	type fields struct {
		Rows *sql.Rows
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SQLRow{
				Rows: tt.fields.Rows,
			}
			if got := r.Next(); got != tt.want {
				t.Errorf("SQLRow.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLRow_Close(t *testing.T) {
	type fields struct {
		Rows *sql.Rows
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := SQLRow{
				Rows: tt.fields.Rows,
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("SQLRow.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
