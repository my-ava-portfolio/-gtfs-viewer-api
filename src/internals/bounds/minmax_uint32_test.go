package bounds

import (
	"testing"
)

func TestBoundsUint32(t *testing.T) {
	values := []uint32{4, 1, 2, 3, 8, 10}

	bounds := GetMinmax_uint32Array(values)

	if  bounds.Min != 1 && bounds.Max != 10 {
		t.Fatalf(`Min and Max expected not found`)
	}
}