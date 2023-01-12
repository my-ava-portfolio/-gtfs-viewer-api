package structures


type Feature interface {
    IsDateValid(date int32) bool
}


type Stop struct {
	Index  			string    	`json:"index"`
	Xcoord        	float64    	`json:"x"`
	Ycoord		 	float64    	`json:"y"`
	StartDate  		int32    	`json:"start_date"`
	EndDate     	int32    	`json:"end_date"`
	RouteType  		string		`json:"route_type"`
	RouteLongName	string    	`json:"route_long_name"`
}


func (stopFeature Stop) IsDateValid(date int32) bool {
	return stopFeature.StartDate <= date && stopFeature.EndDate >= date
}