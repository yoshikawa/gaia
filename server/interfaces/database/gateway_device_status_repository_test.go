package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestGatewayDeviceStatusRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		gatewayDeviceStatus domain.GatewayDeviceStatus
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
			repo := &GatewayDeviceStatusRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.gatewayDeviceStatus)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceStatusRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("GatewayDeviceStatusRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGatewayDeviceStatusRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		wantGatewayDeviceStatus domain.GatewayDeviceStatus
		wantErr                 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &GatewayDeviceStatusRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotGatewayDeviceStatus, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceStatusRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDeviceStatus, tt.wantGatewayDeviceStatus) {
				t.Errorf("GatewayDeviceStatusRepository.FindByID() = %v, want %v", gotGatewayDeviceStatus, tt.wantGatewayDeviceStatus)
			}
		})
	}
}

func TestGatewayDeviceStatusRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name                      string
		fields                    fields
		wantGatewayDeviceStatuses domain.GatewayDeviceStatuses
		wantErr                   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &GatewayDeviceStatusRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotGatewayDeviceStatuses, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceStatusRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDeviceStatuses, tt.wantGatewayDeviceStatuses) {
				t.Errorf("GatewayDeviceStatusRepository.FindAll() = %v, want %v", gotGatewayDeviceStatuses, tt.wantGatewayDeviceStatuses)
			}
		})
	}
}
