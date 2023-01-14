package helpers

import (
	"testing"
	"reflect"
)

func TestGeomBounds(t *testing.T) {
	x := []float32{4.5, 1.8, 2.25, 3.54, 3.78, 3.631, 3.01, 1.5}
	y := []float32{45.63, 101.8, 82.25, 38.54, 377.78, 35.631, 37.01, 17.5}

	bounds := GetBoundsFromXsAndYs(x, y)

	if !reflect.DeepEqual(bounds, [4]float32{1.5, 17.5, 4.5, 377.78}) {
		t.Fatalf(`Geom bounds expected not found`)
	}
}