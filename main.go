package main

import (
	"gtfs_viewer/src/config"
	gtfsStops "gtfs_viewer/src/gtfstops"
	"gtfs_viewer/src/internals/utils"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func selectData(area string) config.FileModel {
	var dataFound config.FileModel

	for _, feature := range data.Files {
		if feature.Title == area {
			dataFound = feature
		}
	}
	return dataFound
}

func movingStopsRoute(context *gin.Context) {
	area := context.Param("area")
    dateParam := context.Query("date")
    
    if dateParam == "" {
		context.String(http.StatusBadRequest, "Param 'date' is missing")
		return
    } else {
		
		dataFound := selectData(area)

		date, _ := strconv.Atoi(dateParam)
		// TODO add error condition check
		stopsFound := gtfsStops.FilterByDate(dataFound.Data, uint32(date))

		context.JSON(http.StatusOK, stopsFound)
 	}	
}


type rangeDataModel struct {
	DataBounds		[4]float32
	StartDate		uint32
	EndDate			uint32
}


func rangeDatesRoute(context *gin.Context) {
	area := context.Param("area")

	dataFound := selectData(area)

	result := rangeDataModel{
		DataBounds: dataFound.Bounds,
		StartDate: dataFound.StartDate,
		EndDate: dataFound.EndDate,
	}
	context.JSON(http.StatusOK, result)
	utils.PrintMemresultUsage()
}



func gtfsGroupRouterRequests(router *gin.Engine) {
	v2 := router.Group("/api/v2/gtfs_builder")

	v2.GET(":area/moving_nodes", movingStopsRoute)
	v2.GET(":area/range_dates", rangeDatesRoute)
	//v2.GET("/route_types", movingStopsRoute)

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://portfolio.amaury-valorge.com"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	  }))
	  
	gtfsGroupRouterRequests(router)

	router.Run(":7001")
}