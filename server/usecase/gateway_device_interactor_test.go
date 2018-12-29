package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestGatewayDeviceInteractor_Add(t *testing.T) {
	type fields struct {
		GatewayDeviceRepository GatewayDeviceRepository
	}
	type args struct {
		u domain.GatewayDevice
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantGatewayDevice domain.GatewayDevice
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &GatewayDeviceInteractor{
				GatewayDeviceRepository: tt.fields.GatewayDeviceRepository,
			}
			gotGatewayDevice, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDevice, tt.wantGatewayDevice) {
				t.Errorf("GatewayDeviceInteractor.Add() = %v, want %v", gotGatewayDevice, tt.wantGatewayDevice)
			}
		})
	}
}

func TestGatewayDeviceInteractor_GatewayDevices(t *testing.T) {
	type fields struct {
		GatewayDeviceRepository GatewayDeviceRepository
	}
	tests := []struct {
		name              string
		fields            fields
		wantGatewayDevice domain.GatewayDevices
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &GatewayDeviceInteractor{
				GatewayDeviceRepository: tt.fields.GatewayDeviceRepository,
			}
			gotGatewayDevice, err := interactor.GatewayDevices()
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceInteractor.GatewayDevices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDevice, tt.wantGatewayDevice) {
				t.Errorf("GatewayDeviceInteractor.GatewayDevices() = %v, want %v", gotGatewayDevice, tt.wantGatewayDevice)
			}
		})
	}
}

func TestGatewayDeviceInteractor_GatewayDeviceByID(t *testing.T) {
	type fields struct {
		GatewayDeviceRepository GatewayDeviceRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantGatewayDevice domain.GatewayDevice
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &GatewayDeviceInteractor{
				GatewayDeviceRepository: tt.fields.GatewayDeviceRepository,
			}
			gotGatewayDevice, err := interactor.GatewayDeviceByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceInteractor.GatewayDeviceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDevice, tt.wantGatewayDevice) {
				t.Errorf("GatewayDeviceInteractor.GatewayDeviceByID() = %v, want %v", gotGatewayDevice, tt.wantGatewayDevice)
			}
		})
	}
}
