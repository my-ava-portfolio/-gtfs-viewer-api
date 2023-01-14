package bounds

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestBoundsFloat32(t *testing.T) {
	values := []float32{4.5, 1.8, 2.25, 3.54, 3.78, 3.631, 3.01, 1.5}

	bounds := GetMinmax_float32Array(values)

    assert.True(t, bounds.Min == 1.5 && bounds.Max == 4.5, `Min and Max expected not found`)

}