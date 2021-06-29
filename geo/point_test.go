package geo

import (
	"math"
	"testing"

	"gopkg.in/validator.v2"
)

const float64EqualityThreshold = 1e-6

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestDistanceBetween(t *testing.T) {
	var fp = Point{-47.621424, -22.696743}
	var sp = Point{-47.646330, -22.702367}
	var distance = fp.DistanceBetween(sp)

	if !almostEqual(distance, 2.804437) {
		t.Errorf("distanceBetween failed, expected %f, got %f", 2.804437, distance)
	}
}

func TestPointValidationSuccess(t *testing.T) {
	p := Point{48, 48}
	if errs := validator.Validate(p); errs != nil {
		t.Errorf("Point validation failed, errors: %s", errs)
	}
}

func TestPointValidationLatitudeError(t *testing.T) {
	p := Point{200, 10}
	if errs := validator.Validate(p); errs == nil {
		t.Errorf("Point validation succeed but must fail, latitude %f is invalid", p.Lat)
	}
}

func TestPointValidationLongitudeError(t *testing.T) {
	p := Point{10, 200}
	if errs := validator.Validate(p); errs == nil {
		t.Errorf("Point validation succeed but must fail, latitude %f is invalid", p.Lat)
	}
}
