package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestSensingDeviceRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		sensingDevice domain.SensingDevice
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
			repo := &SensingDeviceRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.sensingDevice)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("SensingDeviceRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestSensingDeviceRepository_FindByID(t *testing.T) {
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
		wantSensingDevice domain.SensingDevice
		wantErr           bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &SensingDeviceRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotSensingDevice, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDevice, tt.wantSensingDevice) {
				t.Errorf("SensingDeviceRepository.FindByID() = %v, want %v", gotSensingDevice, tt.wantSensingDevice)
			}
		})
	}
}

func TestSensingDeviceRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
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
			repo := &SensingDeviceRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotSensingDevices, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDevices, tt.wantSensingDevices) {
				t.Errorf("SensingDeviceRepository.FindAll() = %v, want %v", gotSensingDevices, tt.wantSensingDevices)
			}
		})
	}
}
