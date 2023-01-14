package helpers

import (
	"gtfs_viewer/src/internals/bounds"
)

func GetBoundsFromXsAndYs(x, y []float32) [4]float32 {
	Xvalues := bounds.GetMinmax_float32Array(x)
	YValues := bounds.GetMinmax_float32Array(y)
	return [4]float32{Xvalues.Min, YValues.Min, Xvalues.Max, YValues.Max}
}
