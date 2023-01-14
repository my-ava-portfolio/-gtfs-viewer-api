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

	result := rangeDataModel{
		DataBounds: dataFound.Bounds,
		StartDate: dataFound.StartDate,
		EndDate: dataFound.EndDate,
	}
	context.JSON(http.StatusOK, result)
	helpers.PrintMemresultUsage()
}

func GtfsGroupRouterRequests(router *gin.Engine) {
	v2 := router.Group("/api/v2/gtfs_builder")

	v2.GET(":area/moving_nodes", movingStopsRoute)
	v2.GET(":area/range_dates", rangeDatesRoute)
	//v2.GET("/route_types", movingStopsRoute)

}