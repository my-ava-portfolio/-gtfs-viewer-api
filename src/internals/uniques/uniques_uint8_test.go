package uniques

import (
	"reflect"
	"testing"
)

func TestUniquesUint8(t *testing.T) {
	values := []uint8{4, 1, 2, 3, 3, 3, 3, 1}

	uniques := Uint8(values)

	if !reflect.DeepEqual(uniques, []uint8{4, 1, 2, 3})  {
		t.Fatalf(`Not equals`)
	}
}