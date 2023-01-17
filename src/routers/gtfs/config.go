package gtfs

type Stop struct {
	//RouteLongName	string    	`json:"route_long_name"`
	Xcoord        	float32    	`json:"x"`
	Ycoord		 	float32    	`json:"y"`
	StartDate  		uint32    	`json:"start_date"`
	EndDate     	uint32    	`json:"end_date"`
	RouteType  		uint8		`json:"route_type"`
}

type RangeDataModel struct {
	StartDate, EndDate		uint32
	DataBounds				[4]float32
}

type FileModel struct {
	Title     			string
	Bounds    			[4]float32
	StartDate, EndDate 	uint32
	Data      			[]Stop
	routeTypes			[]uint8
}

type ConfigModel struct {
	Files []FileModel
}



