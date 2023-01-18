package helpers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeomBounds(t *testing.T) {
	x := []float32{4.5, 1.8, 2.25, 3.54, 3.78, 3.631, 3.01, 1.5}
	y := []float32{45.63, 101.8, 82.25, 38.54, 377.78, 35.631, 37.01, 17.5}

	bounds := GetBoundsFromXsAndYs(x, y)

	assert.True(t, reflect.DeepEqual(bounds, [4]float32{1.5, 17.5, 4.5, 377.78}), `Geom bounds expected not found`)

}

func TestPointIntersectsBounds(t *testing.T) {
	points := [2]float32{2.5, 25.6}
	bounds := []float32{1.5, 17.5, 4.5, 377.78}

	intersects := IsPointIntersectsBounds(points, bounds)

	assert.True(t, intersects, `Point must intersect`)
}

func TestPointNotIntersectsBounds(t *testing.T) {
	points := [2]float32{0, 25.6}
	bounds := []float32{1.5, 17.5, 4.5, 377.78}

	intersects := IsPointIntersectsBounds(points, bounds)

	assert.False(t, intersects, `Point must not intersect`)
}