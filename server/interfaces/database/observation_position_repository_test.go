package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestObservationPositionRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		op domain.ObservationPosition
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
			repo := &ObservationPositionRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationPositionRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("ObservationPositionRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestObservationPositionRepository_FindByID(t *testing.T) {
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
		wantOp  domain.ObservationPosition
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ObservationPositionRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotOp, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationPositionRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOp, tt.wantOp) {
				t.Errorf("ObservationPositionRepository.FindByID() = %v, want %v", gotOp, tt.wantOp)
			}
		})
	}
}

func TestObservationPositionRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name    string
		fields  fields
		wantOps domain.ObservationPositions
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &ObservationPositionRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotOps, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ObservationPositionRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOps, tt.wantOps) {
				t.Errorf("ObservationPositionRepository.FindAll() = %v, want %v", gotOps, tt.wantOps)
			}
		})
	}
}
