package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestSensorRepository_Store(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		sensor domain.Sensor
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
			repo := &SensorRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotId, err := repo.Store(tt.args.sensor)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensorRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("SensorRepository.Store() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestSensorRepository_FindByID(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
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
			repo := &SensorRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotSensor, err := repo.FindByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("SensorRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensor, tt.wantSensor) {
				t.Errorf("SensorRepository.FindByID() = %v, want %v", gotSensor, tt.wantSensor)
			}
		})
	}
}

func TestSensorRepository_FindAll(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
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
			repo := &SensorRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotSensors, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("SensorRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSensors, tt.wantSensors) {
				t.Errorf("SensorRepository.FindAll() = %v, want %v", gotSensors, tt.wantSensors)
			}
		})
	}
}
