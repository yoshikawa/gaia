package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestPlantCategoryInteractor_Add(t *testing.T) {
	type fields struct {
		PlantCategoryRepository PlantCategoryRepository
	}
	type args struct {
		u domain.PlantCategory
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
			interactor := &PlantCategoryInteractor{
				PlantCategoryRepository: tt.fields.PlantCategoryRepository,
			}
			gotPlantCategory, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantCategoryInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlantCategory, tt.wantPlantCategory) {
				t.Errorf("PlantCategoryInteractor.Add() = %v, want %v", gotPlantCategory, tt.wantPlantCategory)
			}
		})
	}
}

func TestPlantCategoryInteractor_PlantCategories(t *testing.T) {
	type fields struct {
		PlantCategoryRepository PlantCategoryRepository
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
			interactor := &PlantCategoryInteractor{
				PlantCategoryRepository: tt.fields.PlantCategoryRepository,
			}
			gotPlantCategories, err := interactor.PlantCategories()
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantCategoryInteractor.PlantCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlantCategories, tt.wantPlantCategories) {
				t.Errorf("PlantCategoryInteractor.PlantCategories() = %v, want %v", gotPlantCategories, tt.wantPlantCategories)
			}
		})
	}
}

func TestPlantCategoryInteractor_PlantCategoryByID(t *testing.T) {
	type fields struct {
		PlantCategoryRepository PlantCategoryRepository
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
			interactor := &PlantCategoryInteractor{
				PlantCategoryRepository: tt.fields.PlantCategoryRepository,
			}
			gotPlantCategory, err := interactor.PlantCategoryByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlantCategoryInteractor.PlantCategoryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPlantCategory, tt.wantPlantCategory) {
				t.Errorf("PlantCategoryInteractor.PlantCategoryByID() = %v, want %v", gotPlantCategory, tt.wantPlantCategory)
			}
		})
	}
}
