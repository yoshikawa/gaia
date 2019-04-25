package domain

import (
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("should return success test validate user", func(t *testing.T) {
		user := &User{
			Name:           "pluslab",
			OrganizationID: 1,
			Email:          "yoshikawa@pluslab.org",
			Password:       "yoshikawataiki",
		}

		err := user.Validate()

		if err != nil {
			t.Error("validate user should return nil of error")
		}
	})

	t.Run("should return error test validate user when field name empty", func(t *testing.T) {
		user := &User{
			Name:     "",
			Email:    "yoshikawa@pluslab.org",
			Password: "yoshikawataiki",
		}

		err := user.Validate()

		if err == nil {
			t.Error("validate user should return error")
		}
	})

	t.Run("should return error test validate user when field password empty", func(t *testing.T) {
		user := &User{
			Name:     "yoshikawataiki",
			Email:    "yoshikawa@pluslab.org",
			Password: "",
		}

		err := user.Validate()

		if err == nil {
			t.Error("validate user should return error")
		}
	})
}
