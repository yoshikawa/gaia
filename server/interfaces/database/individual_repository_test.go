package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestIndividualRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		individual domain.Individual
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
			repo := &IndividualRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.individual)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndividualRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("IndividualRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestIndividualRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantIndividual domain.Individual
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &IndividualRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotIndividual, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndividualRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIndividual, tt.wantIndividual) {
				t.Errorf("IndividualRepository.FindByID() = %v, want %v", gotIndividual, tt.wantIndividual)
			}
		})
	}
}

func TestIndividualRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name            string
		fields          fields
		wantIndividuals domain.Individuals
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &IndividualRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotIndividuals, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("IndividualRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIndividuals, tt.wantIndividuals) {
				t.Errorf("IndividualRepository.FindAll() = %v, want %v", gotIndividuals, tt.wantIndividuals)
			}
		})
	}
}
