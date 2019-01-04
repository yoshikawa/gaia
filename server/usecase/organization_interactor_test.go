package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestOrganizationInteractor_Add(t *testing.T) {
	type fields struct {
		OrganizationRepository OrganizationRepository
	}
	type args struct {
		o domain.Organization
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantOrganization domain.Organization
		wantErr          bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "org_test_1",
		// 	fields: fields{
		// 		nil,
		// 	},
		// 	args: args{
		// 		o: domain.Organization{
		// 			ID:        1,
		// 			Name:      "hoge",
		// 			CreatedAt: time.Now(),
		// 			UpdatedAt: time.Now(),
		// 		},
		// 	},
		// 	wantOrganization: domain.Organization{
		// 		ID:        1,
		// 		Name:      "hoge",
		// 		CreatedAt: time.Now(),
		// 		UpdatedAt: time.Now(),
		// 	},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &OrganizationInteractor{
				OrganizationRepository: tt.fields.OrganizationRepository,
			}
			gotOrganization, err := interactor.Add(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrganization, tt.wantOrganization) {
				t.Errorf("OrganizationInteractor.Add() = %v, want %v", gotOrganization, tt.wantOrganization)
			}
		})
	}
}

func TestOrganizationInteractor_Organizations(t *testing.T) {
	type fields struct {
		OrganizationRepository OrganizationRepository
	}
	tests := []struct {
		name              string
		fields            fields
		wantOrganizations domain.Organizations
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &OrganizationInteractor{
				OrganizationRepository: tt.fields.OrganizationRepository,
			}
			gotOrganizations, err := interactor.Organizations()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationInteractor.Organizations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrganizations, tt.wantOrganizations) {
				t.Errorf("OrganizationInteractor.Organizations() = %v, want %v", gotOrganizations, tt.wantOrganizations)
			}
		})
	}
}

func TestOrganizationInteractor_OrganizationByID(t *testing.T) {
	type fields struct {
		OrganizationRepository OrganizationRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantOrganization domain.Organization
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &OrganizationInteractor{
				OrganizationRepository: tt.fields.OrganizationRepository,
			}
			gotOrganization, err := interactor.OrganizationByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationInteractor.OrganizationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrganization, tt.wantOrganization) {
				t.Errorf("OrganizationInteractor.OrganizationByID() = %v, want %v", gotOrganization, tt.wantOrganization)
			}
		})
	}
}
