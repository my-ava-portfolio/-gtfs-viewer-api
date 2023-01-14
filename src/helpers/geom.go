package helpers

import (
	utils "gtfs_viewer/src/internals"
)

func GetBoundsFromXsAndYs(x, y []float32) [4]float32 {
	Xvalues := utils.GetMinmax_float32Array(x)
	YValues := utils.GetMinmax_float32Array(y)
	return [4]float32{Xvalues.Min, YValues.Min, Xvalues.Max, YValues.Max}
}
