package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestFieldInteractor_Fields(t *testing.T) {
	type fields struct {
		FieldRepository FieldRepository
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
			interactor := &FieldInteractor{
				FieldRepository: tt.fields.FieldRepository,
			}
			gotFields, err := interactor.Fields()
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldInteractor.Fields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFields, tt.wantFields) {
				t.Errorf("FieldInteractor.Fields() = %v, want %v", gotFields, tt.wantFields)
			}
		})
	}
}

func TestFieldInteractor_Add(t *testing.T) {
	type fields struct {
		FieldRepository FieldRepository
	}
	type args struct {
		f domain.Field
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
			interactor := &FieldInteractor{
				FieldRepository: tt.fields.FieldRepository,
			}
			gotField, err := interactor.Add(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotField, tt.wantField) {
				t.Errorf("FieldInteractor.Add() = %v, want %v", gotField, tt.wantField)
			}
		})
	}
}

func TestFieldInteractor_FieldByID(t *testing.T) {
	type fields struct {
		FieldRepository FieldRepository
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
			interactor := &FieldInteractor{
				FieldRepository: tt.fields.FieldRepository,
			}
			gotField, err := interactor.FieldByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("FieldInteractor.FieldByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotField, tt.wantField) {
				t.Errorf("FieldInteractor.FieldByID() = %v, want %v", gotField, tt.wantField)
			}
		})
	}
}
