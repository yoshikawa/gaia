package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestVendorRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		vendor domain.Vendor
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
			repo := &VendorRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.vendor)
			if (err != nil) != tt.wantErr {
				t.Errorf("VendorRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("VendorRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestVendorRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantVendor domain.Vendor
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &VendorRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotVendor, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("VendorRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVendor, tt.wantVendor) {
				t.Errorf("VendorRepository.FindByID() = %v, want %v", gotVendor, tt.wantVendor)
			}
		})
	}
}

func TestVendorRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name        string
		fields      fields
		wantVendors domain.Vendors
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &VendorRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotVendors, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("VendorRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVendors, tt.wantVendors) {
				t.Errorf("VendorRepository.FindAll() = %v, want %v", gotVendors, tt.wantVendors)
			}
		})
	}
}
