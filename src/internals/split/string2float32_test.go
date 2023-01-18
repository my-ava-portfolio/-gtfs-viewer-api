package split

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestBoundsFloat32(t *testing.T) {
	values := "3.5,5,6.5,8,9.65"

	bounds := StringToFloat32(values, ",")

    assert.True(t, len(bounds) == 5, `Elements count expected not found`)

}