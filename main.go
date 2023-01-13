package main

import (
	"gtfs_viewer/src/config"
	gtfsStops "gtfs_viewer/src/gtfstops"
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
		var dataFound []gtfsStops.Stop
		for _, feature := range data.Files {
			if feature.Title == area {
				dataFound = feature.Data
			}
		}
		if dataFound == nil {
			context.String(http.StatusBadRequest, "area not found")
			return
		}

		date, _ := strconv.Atoi(dateParam)
		// TODO add error condition check
		stopsFound := gtfsStops.FilterByDate(dataFound, uint32(date))

		context.JSON(http.StatusOK, stopsFound)
 	}	
}


func gtfsGroupRouterRequests(router *gin.Engine) {
	v2 := router.Group("/api/v2/gtfs_builder")

	v2.GET(":area/moving_nodes", movingStopsRoute)
	//v2.GET("/range_dates", movingStopsRoute)
	//v2.GET("/route_types", movingStopsRoute)
	//v2.GET("/existing_study_areas", movingStopsRoute)

}

var data config.ConfigModel
// invoked before main()
func init() {
    // loads data
    data = config.ParseConfig()

}

func main() {
	// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/24.html
	router := gin.Default()

	gtfsGroupRouterRequests(router)

	router.Run(":8080")
}