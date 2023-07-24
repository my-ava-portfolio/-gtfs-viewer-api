package split

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToUint64WithValidValue(t *testing.T) {
	value := "1"
	converted := StringToUint64(value)

	assert.Equal(t, uint64(1), converted, `string number value not converted to Uint64`)
}

// TODO add test when value is not compatible. Need to rework the method