package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestOrganizationRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		o domain.Organization
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
			repo := &OrganizationRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("OrganizationRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestOrganizationRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
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
			repo := &OrganizationRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotOrganization, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrganization, tt.wantOrganization) {
				t.Errorf("OrganizationRepository.FindByID() = %v, want %v", gotOrganization, tt.wantOrganization)
			}
		})
	}
}

func TestOrganizationRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
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
			repo := &OrganizationRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotOrganizations, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrganizations, tt.wantOrganizations) {
				t.Errorf("OrganizationRepository.FindAll() = %v, want %v", gotOrganizations, tt.wantOrganizations)
			}
		})
	}
}
