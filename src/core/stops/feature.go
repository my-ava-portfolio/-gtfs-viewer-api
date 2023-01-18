package stops

type StopsFeature struct {
	Title              string
	Bounds             [4]float32
	StartDate, EndDate uint32
	Data               []StopItem
	RouteTypes         []uint8
}