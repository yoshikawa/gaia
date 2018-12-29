package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestGatewayDeviceStatusInteractor_Add(t *testing.T) {
	type fields struct {
		GatewayDeviceStatusRepository GatewayDeviceStatusRepository
	}
	type args struct {
		u domain.GatewayDeviceStatus
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
			interactor := &GatewayDeviceStatusInteractor{
				GatewayDeviceStatusRepository: tt.fields.GatewayDeviceStatusRepository,
			}
			gotGatewayDeviceStatus, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceStatusInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDeviceStatus, tt.wantGatewayDeviceStatus) {
				t.Errorf("GatewayDeviceStatusInteractor.Add() = %v, want %v", gotGatewayDeviceStatus, tt.wantGatewayDeviceStatus)
			}
		})
	}
}

func TestGatewayDeviceStatusInteractor_GatewayDeviceStatuses(t *testing.T) {
	type fields struct {
		GatewayDeviceStatusRepository GatewayDeviceStatusRepository
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
			interactor := &GatewayDeviceStatusInteractor{
				GatewayDeviceStatusRepository: tt.fields.GatewayDeviceStatusRepository,
			}
			gotGatewayDeviceStatuses, err := interactor.GatewayDeviceStatuses()
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceStatusInteractor.GatewayDeviceStatuses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDeviceStatuses, tt.wantGatewayDeviceStatuses) {
				t.Errorf("GatewayDeviceStatusInteractor.GatewayDeviceStatuses() = %v, want %v", gotGatewayDeviceStatuses, tt.wantGatewayDeviceStatuses)
			}
		})
	}
}

func TestGatewayDeviceStatusInteractor_GatewayDeviceStatusByID(t *testing.T) {
	type fields struct {
		GatewayDeviceStatusRepository GatewayDeviceStatusRepository
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
			interactor := &GatewayDeviceStatusInteractor{
				GatewayDeviceStatusRepository: tt.fields.GatewayDeviceStatusRepository,
			}
			gotGatewayDeviceStatus, err := interactor.GatewayDeviceStatusByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("GatewayDeviceStatusInteractor.GatewayDeviceStatusByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGatewayDeviceStatus, tt.wantGatewayDeviceStatus) {
				t.Errorf("GatewayDeviceStatusInteractor.GatewayDeviceStatusByID() = %v, want %v", gotGatewayDeviceStatus, tt.wantGatewayDeviceStatus)
			}
		})
	}
}
