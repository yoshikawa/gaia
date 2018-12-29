package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestPlantRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		plant domain.Plant
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
			repo := &PlantRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.plant)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("PlantRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestPlantRepository_FindByID(t *testing.T) {
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
		wantPlant domain.Plant
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PlantRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotPlant, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlant, tt.wantPlant) {
				t.Errorf("PlantRepository.FindByID() = %v, want %v", gotPlant, tt.wantPlant)
			}
		})
	}
}

func TestPlantRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
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
			repo := &PlantRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotPlants, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlants, tt.wantPlants) {
				t.Errorf("PlantRepository.FindAll() = %v, want %v", gotPlants, tt.wantPlants)
			}
		})
	}
}
