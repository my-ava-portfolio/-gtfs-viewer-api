package split

import (
	"strconv"
	"strings"
)

func StringToFloat32(valueToSplit string, delimiter string) []float32 {
	valueSplit := strings.Split(valueToSplit, delimiter)
	valuesFound := make([]float32, len(valueSplit))
	for index, element := range valueSplit {
		element, _ := strconv.ParseFloat(element, 64)
		valuesFound[index] = float32(element)
	}
	return valuesFound
}