package usecase

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestUserInteractor_Add(t *testing.T) {
	type fields struct {
		UserRepository UserRepository
	}
	type args struct {
		u domain.User
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser domain.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &UserInteractor{
				UserRepository: tt.fields.UserRepository,
			}
			gotUser, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("UserInteractor.Add() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserInteractor_UserByID(t *testing.T) {
	type fields struct {
		UserRepository UserRepository
	}
	type args struct {
		identifier int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser domain.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &UserInteractor{
				UserRepository: tt.fields.UserRepository,
			}
			gotUser, err := interactor.UserByID(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.UserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("UserInteractor.UserByID() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestUserInteractor_Users(t *testing.T) {
	type fields struct {
		UserRepository UserRepository
	}
	tests := []struct {
		name      string
		fields    fields
		wantUsers domain.Users
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &UserInteractor{
				UserRepository: tt.fields.UserRepository,
			}
			gotUsers, err := interactor.Users()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.Users() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUsers, tt.wantUsers) {
				t.Errorf("UserInteractor.Users() = %v, want %v", gotUsers, tt.wantUsers)
			}
		})
	}
}
