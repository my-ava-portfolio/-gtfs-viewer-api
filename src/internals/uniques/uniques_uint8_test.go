package uniques

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquesUint8(t *testing.T) {
	values := []uint8{4, 1, 2, 3, 3, 3, 3, 1}

	uniques := Uint8(values)

	assert.True(t, reflect.DeepEqual(uniques, []uint8{4, 1, 2, 3}), `Min and Max expected not found`)

}