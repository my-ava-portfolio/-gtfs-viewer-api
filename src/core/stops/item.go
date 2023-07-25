package stops

import "gtfs_viewer/src/helpers"

type StopItem struct {
	Xcoord    float32 `json:"x"`
	Ycoord    float32 `json:"y"`
	Index     uint32  `json:"index"`
	StartDate uint32  `json:"start_date"`
	EndDate   uint32  `json:"end_date"`
	RouteType uint8   `json:"route_type"`
	RouteId   uint8   `json:"route_id"`
}
func (s *StopItem) IsDateValid(date uint32) bool {
	return s.StartDate <= date && s.EndDate >= date
}
func (s *StopItem) IntersectsBounds(bounds []float32) bool {
	return helpers.IsPointIntersectsBounds([2]float32{s.Xcoord, s.Ycoord}, bounds)
}

type StopRouteItem struct {
	RouteId   		uint32 `json:"route_id"`
	RouteLongName   string `json:"route_long_name"`
}