package gtfs

import (
	"gtfs_viewer/src/helpers"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func movingStopsRoute(context *gin.Context) {
	area := context.Param("area")
    dateParam := context.Query("date")
    
    if dateParam == "" {
		context.String(http.StatusBadRequest, "Param 'date' is missing")
		return
    } else {
		
		dataFound := SelectData(area)

		date, _ := strconv.Atoi(dateParam)
		// TODO add error condition check
		stopsFound := FilterByDate(dataFound.Data, uint32(date))

		context.JSON(http.StatusOK, stopsFound)
 	}	
}

func rangeDatesRoute(context *gin.Context) {
	area := context.Param("area")

	dataFound := SelectData(area)

	result := RangeDataModel{
		DataBounds: dataFound.Bounds,
		StartDate: dataFound.StartDate,
		EndDate: dataFound.EndDate,
	}
	context.JSON(http.StatusOK, result)
	helpers.PrintMemresultUsage()
}

func transportTypeRoute(context *gin.Context) {
	area := context.Param("area")

	dataFound := SelectData(area)
	context.JSON(http.StatusOK, dataFound.routeTypes)
}

func availableAreasRoute(context *gin.Context) {
	var availableAreas []string
	for _, feature := range GtfsInputData.Files {
		availableAreas = append(availableAreas, feature.Title)
	}
	context.JSON(http.StatusOK, availableAreas)
}

func GtfsGroupRouterRequests(router *gin.Engine) {
	group := router.Group("/api/v2/gtfs_builder")

	group.GET(":area/moving_nodes", movingStopsRoute)
	group.GET(":area/range_dates", rangeDatesRoute)
	group.GET(":area/route_types", transportTypeRoute)
	group.GET("/existing_study_areas", availableAreasRoute)

	
}