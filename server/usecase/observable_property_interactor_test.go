package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestObservablePropertyInteractor_Add(t *testing.T) {
	type fields struct {
		ObservablePropertyRepository ObservablePropertyRepository
	}
	type args struct {
		u domain.ObservableProperty
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
			interactor := &ObservablePropertyInteractor{
				ObservablePropertyRepository: tt.fields.ObservablePropertyRepository,
			}
			gotObservableProperty, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservablePropertyInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservableProperty, tt.wantObservableProperty) {
				t.Errorf("ObservablePropertyInteractor.Add() = %v, want %v", gotObservableProperty, tt.wantObservableProperty)
			}
		})
	}
}

func TestObservablePropertyInteractor_ObservableProperties(t *testing.T) {
	type fields struct {
		ObservablePropertyRepository ObservablePropertyRepository
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
			interactor := &ObservablePropertyInteractor{
				ObservablePropertyRepository: tt.fields.ObservablePropertyRepository,
			}
			gotObservableProperties, err := interactor.ObservableProperties()
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservablePropertyInteractor.ObservableProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservableProperties, tt.wantObservableProperties) {
				t.Errorf("ObservablePropertyInteractor.ObservableProperties() = %v, want %v", gotObservableProperties, tt.wantObservableProperties)
			}
		})
	}
}

func TestObservablePropertyInteractor_ObservablePropertyByID(t *testing.T) {
	type fields struct {
		ObservablePropertyRepository ObservablePropertyRepository
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
			interactor := &ObservablePropertyInteractor{
				ObservablePropertyRepository: tt.fields.ObservablePropertyRepository,
			}
			gotObservableProperty, err := interactor.ObservablePropertyByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservablePropertyInteractor.ObservablePropertyByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservableProperty, tt.wantObservableProperty) {
				t.Errorf("ObservablePropertyInteractor.ObservablePropertyByID() = %v, want %v", gotObservableProperty, tt.wantObservableProperty)
			}
		})
	}
}
