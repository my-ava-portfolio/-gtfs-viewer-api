package bounds

import (
	"testing"
)

func TestBoundsFloat32(t *testing.T) {
	values := []float32{4.5, 1.8, 2.25, 3.54, 3.78, 3.631, 3.01, 1.5}

	bounds := GetMinmax_float32Array(values)

	if  bounds.Min != 1.5 && bounds.Max != 4.5 {
		t.Fatalf(`Min and Max expected not found`)
	}
}