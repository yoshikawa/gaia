package controller

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/interfaces/database"
	"github.com/Pluslab/gaia/server/usecase"
)

func TestNewGatewayDeviceStatusController(t *testing.T) {
	type args struct {
		SQLHandler database.SQLHandler
	}
	tests := []struct {
		name string
		args args
		want *GatewayDeviceStatusController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGatewayDeviceStatusController(tt.args.SQLHandler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGatewayDeviceStatusController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGatewayDeviceStatusController_Create(t *testing.T) {
	type fields struct {
		Interactor usecase.GatewayDeviceStatusInteractor
	}
	type args struct {
		c Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &GatewayDeviceStatusController{
				Interactor: tt.fields.Interactor,
			}
			controller.Create(tt.args.c)
		})
	}
}

func TestGatewayDeviceStatusController_Index(t *testing.T) {
	type fields struct {
		Interactor usecase.GatewayDeviceStatusInteractor
	}
	type args struct {
		c Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &GatewayDeviceStatusController{
				Interactor: tt.fields.Interactor,
			}
			controller.Index(tt.args.c)
		})
	}
}

func TestGatewayDeviceStatusController_Show(t *testing.T) {
	type fields struct {
		Interactor usecase.GatewayDeviceStatusInteractor
	}
	type args struct {
		c Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &GatewayDeviceStatusController{
				Interactor: tt.fields.Interactor,
			}
			controller.Show(tt.args.c)
		})
	}
}
