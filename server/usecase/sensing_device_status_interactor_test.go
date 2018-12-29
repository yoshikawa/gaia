package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestSensingDeviceStatusInteractor_Add(t *testing.T) {
	type fields struct {
		SensingDeviceStatusRepository SensingDeviceStatusRepository
	}
	type args struct {
		u domain.SensingDeviceStatus
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
			interactor := &SensingDeviceStatusInteractor{
				SensingDeviceStatusRepository: tt.fields.SensingDeviceStatusRepository,
			}
			gotSensingDeviceStatus, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceStatusInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDeviceStatus, tt.wantSensingDeviceStatus) {
				t.Errorf("SensingDeviceStatusInteractor.Add() = %v, want %v", gotSensingDeviceStatus, tt.wantSensingDeviceStatus)
			}
		})
	}
}

func TestSensingDeviceStatusInteractor_SensingDeviceStatuses(t *testing.T) {
	type fields struct {
		SensingDeviceStatusRepository SensingDeviceStatusRepository
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
			interactor := &SensingDeviceStatusInteractor{
				SensingDeviceStatusRepository: tt.fields.SensingDeviceStatusRepository,
			}
			gotSensingDeviceStatuses, err := interactor.SensingDeviceStatuses()
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceStatusInteractor.SensingDeviceStatuses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDeviceStatuses, tt.wantSensingDeviceStatuses) {
				t.Errorf("SensingDeviceStatusInteractor.SensingDeviceStatuses() = %v, want %v", gotSensingDeviceStatuses, tt.wantSensingDeviceStatuses)
			}
		})
	}
}

func TestSensingDeviceStatusInteractor_SensingDeviceStatusByID(t *testing.T) {
	type fields struct {
		SensingDeviceStatusRepository SensingDeviceStatusRepository
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
			interactor := &SensingDeviceStatusInteractor{
				SensingDeviceStatusRepository: tt.fields.SensingDeviceStatusRepository,
			}
			gotSensingDeviceStatus, err := interactor.SensingDeviceStatusByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensingDeviceStatusInteractor.SensingDeviceStatusByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensingDeviceStatus, tt.wantSensingDeviceStatus) {
				t.Errorf("SensingDeviceStatusInteractor.SensingDeviceStatusByID() = %v, want %v", gotSensingDeviceStatus, tt.wantSensingDeviceStatus)
			}
		})
	}
}
