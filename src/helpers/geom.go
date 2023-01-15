package helpers

import (
	"gtfs_viewer/src/internals/bounds"
)

func GetBoundsFromXsAndYs(x, y []float32) [4]float32 {
	Xvalues := bounds.GetMinmax_float32Array(x)
	YValues := bounds.GetMinmax_float32Array(y)
	return [4]float32{Xvalues.Min, YValues.Min, Xvalues.Max, YValues.Max}
}


func IsPointIntersectsBounds(point [2]float32, bounds[4]float32) bool {
	
	if (point[0] >= bounds[0] && point[0] <= bounds[2]) && (point[1] >= bounds[1] && point[1] <= bounds[3]) {
		return true
	}
	return false
}