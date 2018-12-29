package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestSensingDeviceInteractor_Add(t *testing.T) {
	type fields struct {
		SensingDeviceRepository SensingDeviceRepository
	}
	type args struct {
		u domain.SensingDevice
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantSensingDevice domain.SensingDevice
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &SensingDeviceInteractor{
				SensingDeviceRepository: tt.fields.SensingDeviceRepository,
			}
			gotSensingDevice, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDevice, tt.wantSensingDevice) {
				t.Errorf("SensingDeviceInteractor.Add() = %v, want %v", gotSensingDevice, tt.wantSensingDevice)
			}
		})
	}
}

func TestSensingDeviceInteractor_SensingDevices(t *testing.T) {
	type fields struct {
		SensingDeviceRepository SensingDeviceRepository
	}
	tests := []struct {
		name               string
		fields             fields
		wantSensingDevices domain.SensingDevices
		wantErr            bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &SensingDeviceInteractor{
				SensingDeviceRepository: tt.fields.SensingDeviceRepository,
			}
			gotSensingDevices, err := interactor.SensingDevices()
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceInteractor.SensingDevices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDevices, tt.wantSensingDevices) {
				t.Errorf("SensingDeviceInteractor.SensingDevices() = %v, want %v", gotSensingDevices, tt.wantSensingDevices)
			}
		})
	}
}

func TestSensingDeviceInteractor_SensingDeviceByID(t *testing.T) {
	type fields struct {
		SensingDeviceRepository SensingDeviceRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantSensingDevice domain.SensingDevice
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &SensingDeviceInteractor{
				SensingDeviceRepository: tt.fields.SensingDeviceRepository,
			}
			gotSensingDevice, err := interactor.SensingDeviceByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceInteractor.SensingDeviceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDevice, tt.wantSensingDevice) {
				t.Errorf("SensingDeviceInteractor.SensingDeviceByID() = %v, want %v", gotSensingDevice, tt.wantSensingDevice)
			}
		})
	}
}
