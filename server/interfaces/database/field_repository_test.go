package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestFieldRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		f domain.Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &FieldRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("FieldRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestFieldRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantField domain.Field
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &FieldRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotField, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotField, tt.wantField) {
				t.Errorf("FieldRepository.FindByID() = %v, want %v", gotField, tt.wantField)
			}
		})
	}
}

func TestFieldRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name       string
		fields     fields
		wantFields domain.Fields
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &FieldRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotFields, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFields, tt.wantFields) {
				t.Errorf("FieldRepository.FindAll() = %v, want %v", gotFields, tt.wantFields)
			}
		})
	}
}
