package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestPlantCategoryRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		plantCategory domain.PlantCategory
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
			repo := &PlantCategoryRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.plantCategory)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantCategoryRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("PlantCategoryRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestPlantCategoryRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantPlantCategory domain.PlantCategory
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PlantCategoryRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotPlantCategory, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantCategoryRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlantCategory, tt.wantPlantCategory) {
				t.Errorf("PlantCategoryRepository.FindByID() = %v, want %v", gotPlantCategory, tt.wantPlantCategory)
			}
		})
	}
}

func TestPlantCategoryRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name                string
		fields              fields
		wantPlantCategories domain.PlantCategories
		wantErr             bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PlantCategoryRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotPlantCategories, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantCategoryRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlantCategories, tt.wantPlantCategories) {
				t.Errorf("PlantCategoryRepository.FindAll() = %v, want %v", gotPlantCategories, tt.wantPlantCategories)
			}
		})
	}
}
