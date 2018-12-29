package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestSensingDeviceStatusRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		sensingDeviceStatus domain.SensingDeviceStatus
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
			repo := &SensingDeviceStatusRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.sensingDeviceStatus)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceStatusRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("SensingDeviceStatusRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestSensingDeviceStatusRepository_FindByID(t *testing.T) {
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
		wantSensingDeviceStatus domain.SensingDeviceStatus
		wantErr                 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &SensingDeviceStatusRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotSensingDeviceStatus, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceStatusRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDeviceStatus, tt.wantSensingDeviceStatus) {
				t.Errorf("SensingDeviceStatusRepository.FindByID() = %v, want %v", gotSensingDeviceStatus, tt.wantSensingDeviceStatus)
			}
		})
	}
}

func TestSensingDeviceStatusRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	tests := []struct {
		name                      string
		fields                    fields
		wantSensingDeviceStatuses domain.SensingDeviceStatuses
		wantErr                   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &SensingDeviceStatusRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotSensingDeviceStatuses, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceStatusRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDeviceStatuses, tt.wantSensingDeviceStatuses) {
				t.Errorf("SensingDeviceStatusRepository.FindAll() = %v, want %v", gotSensingDeviceStatuses, tt.wantSensingDeviceStatuses)
			}
		})
	}
}
