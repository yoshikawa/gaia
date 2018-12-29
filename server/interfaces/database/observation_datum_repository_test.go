package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestObservationDatumRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		od domain.ObservationDatum
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
			repo := &ObservationDatumRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.od)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationDatumRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("ObservationDatumRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestObservationDatumRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantOd  domain.ObservationDatum
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ObservationDatumRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotOd, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationDatumRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOd, tt.wantOd) {
				t.Errorf("ObservationDatumRepository.FindByID() = %v, want %v", gotOd, tt.wantOd)
			}
		})
	}
}

func TestObservationDatumRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name    string
		fields  fields
		wantOd  domain.ObservationData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ObservationDatumRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotOd, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationDatumRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOd, tt.wantOd) {
				t.Errorf("ObservationDatumRepository.FindAll() = %v, want %v", gotOd, tt.wantOd)
			}
		})
	}
}
