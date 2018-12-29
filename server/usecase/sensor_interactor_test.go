package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestSensorInteractor_Add(t *testing.T) {
	type fields struct {
		SensorRepository SensorRepository
	}
	type args struct {
		u domain.Sensor
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantSensor domain.Sensor
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &SensorInteractor{
				SensorRepository: tt.fields.SensorRepository,
			}
			gotSensor, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensorInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensor, tt.wantSensor) {
				t.Errorf("SensorInteractor.Add() = %v, want %v", gotSensor, tt.wantSensor)
			}
		})
	}
}

func TestSensorInteractor_Sensors(t *testing.T) {
	type fields struct {
		SensorRepository SensorRepository
	}
	tests := []struct {
		name        string
		fields      fields
		wantSensors domain.Sensors
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &SensorInteractor{
				SensorRepository: tt.fields.SensorRepository,
			}
			gotSensors, err := interactor.Sensors()
			if (err != nil) != tt.wantErr {
				t.Errorf("SensorInteractor.Sensors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensors, tt.wantSensors) {
				t.Errorf("SensorInteractor.Sensors() = %v, want %v", gotSensors, tt.wantSensors)
			}
		})
	}
}

func TestSensorInteractor_SensorByID(t *testing.T) {
	type fields struct {
		SensorRepository SensorRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantSensor domain.Sensor
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &SensorInteractor{
				SensorRepository: tt.fields.SensorRepository,
			}
			gotSensor, err := interactor.SensorByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensorInteractor.SensorByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensor, tt.wantSensor) {
				t.Errorf("SensorInteractor.SensorByID() = %v, want %v", gotSensor, tt.wantSensor)
			}
		})
	}
}
