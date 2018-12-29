package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestPlantInteractor_Add(t *testing.T) {
	type fields struct {
		PlantRepository PlantRepository
	}
	type args struct {
		u domain.Plant
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantPlant domain.Plant
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &PlantInteractor{
				PlantRepository: tt.fields.PlantRepository,
			}
			gotPlant, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlant, tt.wantPlant) {
				t.Errorf("PlantInteractor.Add() = %v, want %v", gotPlant, tt.wantPlant)
			}
		})
	}
}

func TestPlantInteractor_Plants(t *testing.T) {
	type fields struct {
		PlantRepository PlantRepository
	}
	tests := []struct {
		name       string
		fields     fields
		wantPlants domain.Plants
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &PlantInteractor{
				PlantRepository: tt.fields.PlantRepository,
			}
			gotPlants, err := interactor.Plants()
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantInteractor.Plants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlants, tt.wantPlants) {
				t.Errorf("PlantInteractor.Plants() = %v, want %v", gotPlants, tt.wantPlants)
			}
		})
	}
}

func TestPlantInteractor_PlantByID(t *testing.T) {
	type fields struct {
		PlantRepository PlantRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantPlant domain.Plant
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &PlantInteractor{
				PlantRepository: tt.fields.PlantRepository,
			}
			gotPlant, err := interactor.PlantByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantInteractor.PlantByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlant, tt.wantPlant) {
				t.Errorf("PlantInteractor.PlantByID() = %v, want %v", gotPlant, tt.wantPlant)
			}
		})
	}
}
