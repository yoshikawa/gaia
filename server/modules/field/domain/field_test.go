package domain

import (
	"testing"

	"github.com/paulmach/orb"
)

func TestField(t *testing.T) {
	t.Run("should return success test validate field", func(t *testing.T) {
		poly := orb.Polygon{
			{
				{-122.4163816, 37.7792782},
				{-122.4162786, 37.7787626},
				{-122.4151027, 37.7789118},
				{-122.4152143, 37.7794274},
				{-122.4163816, 37.7792782},
			},
		}
		field := &Field{
			Name: "pluslab",
			Area: poly,
		}

		err := field.Validate()

		if err != nil {
			t.Error("validate field should return nil of error")
		}
	})

	t.Run("should return error test validate field when field name empty", func(t *testing.T) {
		poly := orb.Polygon{
			{
				{-122.4163816, 37.7792782},
				{-122.4162786, 37.7787626},
				{-122.4151027, 37.7789118},
				{-122.4152143, 37.7794274},
				{-122.4163816, 37.7792782},
			},
		}

		field := &Field{
			Name: "",
			Area: poly,
		}

		err := field.Validate()

		if err == nil {
			t.Error("validate field should return error")
		}
	})

	t.Run("should return error test validate field when field area empty", func(t *testing.T) {
		field := &Field{
			Name: "hoge",
		}

		err := field.Validate()

		if err == nil {
			t.Error("validate field should return error")
		}
	})
}
