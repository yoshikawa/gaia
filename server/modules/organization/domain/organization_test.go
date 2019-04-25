package domain

import (
	"testing"
)

func TestOrganization(t *testing.T) {
	t.Run("should return success test validate organization", func(t *testing.T) {
		organization := &Organization{
			Name: "pluslab",
		}

		err := organization.Validate()

		if err != nil {
			t.Error("validate organization should return nil of error")
		}
	})

	t.Run("should return error test validate organization when field name empty", func(t *testing.T) {
		organization := &Organization{
			Name: "",
		}

		err := organization.Validate()

		if err == nil {
			t.Error("validate organization should return error")
		}
	})
}
