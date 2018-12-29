package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestGatewayDeviceRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		gatewayDevice domain.GatewayDevice
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
			repo := &GatewayDeviceRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.gatewayDevice)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("GatewayDeviceRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGatewayDeviceRepository_FindByID(t *testing.T) {
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
		wantGatewayDevice domain.GatewayDevice
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &GatewayDeviceRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotGatewayDevice, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDevice, tt.wantGatewayDevice) {
				t.Errorf("GatewayDeviceRepository.FindByID() = %v, want %v", gotGatewayDevice, tt.wantGatewayDevice)
			}
		})
	}
}

func TestGatewayDeviceRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name               string
		fields             fields
		wantGatewayDevices domain.GatewayDevices
		wantErr            bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &GatewayDeviceRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotGatewayDevices, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDevices, tt.wantGatewayDevices) {
				t.Errorf("GatewayDeviceRepository.FindAll() = %v, want %v", gotGatewayDevices, tt.wantGatewayDevices)
			}
		})
	}
}
