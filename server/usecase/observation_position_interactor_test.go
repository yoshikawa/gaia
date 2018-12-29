package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestObservationPositionInteractor_Add(t *testing.T) {
	type fields struct {
		ObservationPositionRepository ObservationPositionRepository
	}
	type args struct {
		u domain.ObservationPosition
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		wantObservationPosition domain.ObservationPosition
		wantErr                 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &ObservationPositionInteractor{
				ObservationPositionRepository: tt.fields.ObservationPositionRepository,
			}
			gotObservationPosition, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationPositionInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservationPosition, tt.wantObservationPosition) {
				t.Errorf("ObservationPositionInteractor.Add() = %v, want %v", gotObservationPosition, tt.wantObservationPosition)
			}
		})
	}
}

func TestObservationPositionInteractor_ObservationPositions(t *testing.T) {
	type fields struct {
		ObservationPositionRepository ObservationPositionRepository
	}
	tests := []struct {
		name                     string
		fields                   fields
		wantObservationPositions domain.ObservationPositions
		wantErr                  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &ObservationPositionInteractor{
				ObservationPositionRepository: tt.fields.ObservationPositionRepository,
			}
			gotObservationPositions, err := interactor.ObservationPositions()
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationPositionInteractor.ObservationPositions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservationPositions, tt.wantObservationPositions) {
				t.Errorf("ObservationPositionInteractor.ObservationPositions() = %v, want %v", gotObservationPositions, tt.wantObservationPositions)
			}
		})
	}
}

func TestObservationPositionInteractor_ObservationPositionByID(t *testing.T) {
	type fields struct {
		ObservationPositionRepository ObservationPositionRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		wantObservationPosition domain.ObservationPosition
		wantErr                 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &ObservationPositionInteractor{
				ObservationPositionRepository: tt.fields.ObservationPositionRepository,
			}
			gotObservationPosition, err := interactor.ObservationPositionByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationPositionInteractor.ObservationPositionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObservationPosition, tt.wantObservationPosition) {
				t.Errorf("ObservationPositionInteractor.ObservationPositionByID() = %v, want %v", gotObservationPosition, tt.wantObservationPosition)
			}
		})
	}
}
