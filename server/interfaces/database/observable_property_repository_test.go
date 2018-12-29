package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestObservablePropertyRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		observableProperty domain.ObservableProperty
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
			repo := &ObservablePropertyRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.observableProperty)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservablePropertyRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("ObservablePropertyRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestObservablePropertyRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		wantObservableProperty domain.ObservableProperty
		wantErr                bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ObservablePropertyRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotObservableProperty, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservablePropertyRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservableProperty, tt.wantObservableProperty) {
				t.Errorf("ObservablePropertyRepository.FindByID() = %v, want %v", gotObservableProperty, tt.wantObservableProperty)
			}
		})
	}
}

func TestObservablePropertyRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name                     string
		fields                   fields
		wantObservableProperties domain.ObservableProperties
		wantErr                  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ObservablePropertyRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotObservableProperties, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservablePropertyRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservableProperties, tt.wantObservableProperties) {
				t.Errorf("ObservablePropertyRepository.FindAll() = %v, want %v", gotObservableProperties, tt.wantObservableProperties)
			}
		})
	}
}
