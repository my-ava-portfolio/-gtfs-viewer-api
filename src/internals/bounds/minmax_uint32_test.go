package bounds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoundsUint32(t *testing.T) {
	values := []uint32{4, 1, 2, 3, 8, 10}

	bounds := GetMinmax_uint32Array(values)

	assert.True(t, bounds.Min == 1 && bounds.Max == 10, `Min and Max expected not found`)

}