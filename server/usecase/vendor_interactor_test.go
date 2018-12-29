package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestVendorInteractor_Add(t *testing.T) {
	type fields struct {
		VendorRepository VendorRepository
	}
	type args struct {
		u domain.Vendor
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
			interactor := &VendorInteractor{
				VendorRepository: tt.fields.VendorRepository,
			}
			gotVendor, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("VendorInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVendor, tt.wantVendor) {
				t.Errorf("VendorInteractor.Add() = %v, want %v", gotVendor, tt.wantVendor)
			}
		})
	}
}

func TestVendorInteractor_Vendors(t *testing.T) {
	type fields struct {
		VendorRepository VendorRepository
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
			interactor := &VendorInteractor{
				VendorRepository: tt.fields.VendorRepository,
			}
			gotVendors, err := interactor.Vendors()
			if (err != nil) != tt.wantErr {
				t.Errorf("VendorInteractor.Vendors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVendors, tt.wantVendors) {
				t.Errorf("VendorInteractor.Vendors() = %v, want %v", gotVendors, tt.wantVendors)
			}
		})
	}
}

func TestVendorInteractor_VendorByID(t *testing.T) {
	type fields struct {
		VendorRepository VendorRepository
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
			interactor := &VendorInteractor{
				VendorRepository: tt.fields.VendorRepository,
			}
			gotVendor, err := interactor.VendorByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("VendorInteractor.VendorByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVendor, tt.wantVendor) {
				t.Errorf("VendorInteractor.VendorByID() = %v, want %v", gotVendor, tt.wantVendor)
			}
		})
	}
}
