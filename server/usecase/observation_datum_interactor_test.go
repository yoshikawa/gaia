package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestObservationDatumInteractor_Add(t *testing.T) {
	type fields struct {
		ObservationDatumRepository ObservationDatumRepository
	}
	type args struct {
		u domain.ObservationDatum
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantObservationDatum domain.ObservationDatum
		wantErr              bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &ObservationDatumInteractor{
				ObservationDatumRepository: tt.fields.ObservationDatumRepository,
			}
			gotObservationDatum, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationDatumInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservationDatum, tt.wantObservationDatum) {
				t.Errorf("ObservationDatumInteractor.Add() = %v, want %v", gotObservationDatum, tt.wantObservationDatum)
			}
		})
	}
}

func TestObservationDatumInteractor_ObservationData(t *testing.T) {
	type fields struct {
		ObservationDatumRepository ObservationDatumRepository
	}
	tests := []struct {
		name                 string
		fields               fields
		wantObservationDatum domain.ObservationData
		wantErr              bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &ObservationDatumInteractor{
				ObservationDatumRepository: tt.fields.ObservationDatumRepository,
			}
			gotObservationDatum, err := interactor.ObservationData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationDatumInteractor.ObservationData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservationDatum, tt.wantObservationDatum) {
				t.Errorf("ObservationDatumInteractor.ObservationData() = %v, want %v", gotObservationDatum, tt.wantObservationDatum)
			}
		})
	}
}

func TestObservationDatumInteractor_ObservationDatumByID(t *testing.T) {
	type fields struct {
		ObservationDatumRepository ObservationDatumRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantObservationDatum domain.ObservationDatum
		wantErr              bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &ObservationDatumInteractor{
				ObservationDatumRepository: tt.fields.ObservationDatumRepository,
			}
			gotObservationDatum, err := interactor.ObservationDatumByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationDatumInteractor.ObservationDatumByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservationDatum, tt.wantObservationDatum) {
				t.Errorf("ObservationDatumInteractor.ObservationDatumByID() = %v, want %v", gotObservationDatum, tt.wantObservationDatum)
			}
		})
	}
}
