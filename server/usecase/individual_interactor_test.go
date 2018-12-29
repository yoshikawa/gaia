package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestIndividualInteractor_Add(t *testing.T) {
	type fields struct {
		IndividualRepository IndividualRepository
	}
	type args struct {
		u domain.Individual
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
			interactor := &IndividualInteractor{
				IndividualRepository: tt.fields.IndividualRepository,
			}
			gotIndividual, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndividualInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIndividual, tt.wantIndividual) {
				t.Errorf("IndividualInteractor.Add() = %v, want %v", gotIndividual, tt.wantIndividual)
			}
		})
	}
}

func TestIndividualInteractor_Individuals(t *testing.T) {
	type fields struct {
		IndividualRepository IndividualRepository
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
			interactor := &IndividualInteractor{
				IndividualRepository: tt.fields.IndividualRepository,
			}
			gotIndividuals, err := interactor.Individuals()
			if (err != nil) != tt.wantErr {
				t.Errorf("IndividualInteractor.Individuals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIndividuals, tt.wantIndividuals) {
				t.Errorf("IndividualInteractor.Individuals() = %v, want %v", gotIndividuals, tt.wantIndividuals)
			}
		})
	}
}

func TestIndividualInteractor_IndividualByID(t *testing.T) {
	type fields struct {
		IndividualRepository IndividualRepository
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
			interactor := &IndividualInteractor{
				IndividualRepository: tt.fields.IndividualRepository,
			}
			gotIndividual, err := interactor.IndividualByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndividualInteractor.IndividualByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIndividual, tt.wantIndividual) {
				t.Errorf("IndividualInteractor.IndividualByID() = %v, want %v", gotIndividual, tt.wantIndividual)
			}
		})
	}
}
