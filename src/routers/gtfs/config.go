package gtfs

type Stop struct {
	RouteLongName	string    	`json:"route_long_name"`
	Xcoord        	float32    	`json:"x"`
	Ycoord		 	float32    	`json:"y"`
	StartDate  		uint32    	`json:"start_date"`
	EndDate     	uint32    	`json:"end_date"`
	RouteType  		uint8		`json:"route_type"`
}

type rangeDataModel struct {
	DataBounds		[4]float32
	StartDate		uint32
	EndDate			uint32
}

type FileModel struct {
	Title     string
	Data      []Stop
	Bounds    [4]float32
	StartDate uint32
	EndDate   uint32
}

type ConfigModel struct {
	Files []FileModel
}

func ParseConfig() ConfigModel {
	filesFound := getData("data/", "_data.json")

	config := ConfigModel{
		Files: filesFound,
	}
	return config
}

