package gtfs

import (
	gtfsStops "gtfs_viewer/src/core/stops"
	"gtfs_viewer/src/helpers"
	"gtfs_viewer/src/internals/split"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var GlobalGtfsStopData gtfsStops.StopsContainer


func movingStopsRoute(context *gin.Context) {
	area := context.Param("area")

	dateParam := context.Query("date")
	bounds := context.Query("bounds")

	if dateParam == "" {
		context.String(http.StatusBadRequest, "Param 'date' is missing")
		return
	} else {

		date, err := strconv.ParseUint(dateParam, 10, 32)
		if err != nil {
			context.String(http.StatusBadRequest, "Param 'date' not relevant")
			return
		}
		boundsValues := split.StringToFloat32(bounds, ",")

		stopsFound := GlobalGtfsStopData.GetStopsFilteredData(area, uint32(date), boundsValues)
		context.JSON(http.StatusOK, stopsFound)
	}
}

func rangeDatesRoute(context *gin.Context) {
	area := context.Param("area")
	context.JSON(http.StatusOK, 
				 GlobalGtfsStopData.GetRangesData(area))
}

func transportTypeRoute(context *gin.Context) {
	area := context.Param("area")
	context.JSON(http.StatusOK, 
				 GlobalGtfsStopData.GetAreaRouteTypes(area))
}

func availableAreasRoute(context *gin.Context) {
	//availableAreas := GetAreas()
	context.JSON(http.StatusOK, 
				 GlobalGtfsStopData.GetAreas())
}

func GtfsGroupRouterHandler(dataPath string, router *gin.Engine) {

	// get data and set the data global var about GtfsStopsData
	gtfsStopSuffix := "_gtfsData.json"
	GlobalGtfsStopData = gtfsStops.GetData(dataPath, gtfsStopSuffix)

	group := router.Group("/api/v2")

	group.GET(":area/moving_nodes", movingStopsRoute)
	group.GET(":area/range_dates", rangeDatesRoute)
	group.GET(":area/route_types", transportTypeRoute)
	group.GET("/existing_study_areas", availableAreasRoute)
	helpers.PrintMemresultUsage()

}
