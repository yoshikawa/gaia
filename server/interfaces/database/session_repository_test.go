package database

import (
	"reflect"
	"testing"

	"github.com/Pluslab/gaia/server/domain"
)

func TestSessionRepository_Login(t *testing.T) {
	type fields struct {
		SQLHandler SQLHandler
	}
	type args struct {
		inputEmail string
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
			repo := &SessionRepository{
				SQLHandler: tt.fields.SQLHandler,
			}
			gotUser, err := repo.Login(tt.args.inputEmail)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionRepository.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("SessionRepository.Login() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
