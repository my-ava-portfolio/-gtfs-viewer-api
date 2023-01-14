package gtfs

import (
	"gtfs_viewer/src/helpers"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)




func selectData(area string) FileModel {
	var dataFound FileModel

	for _, feature := range GtfsInputData.Files {
		if feature.Title == area {
			dataFound = feature
		}
	}
	return dataFound
}

func MovingStopsRoute(context *gin.Context) {
	area := context.Param("area")
    dateParam := context.Query("date")
    
    if dateParam == "" {
		context.String(http.StatusBadRequest, "Param 'date' is missing")
		return
    } else {
		
		dataFound := selectData(area)

		date, _ := strconv.Atoi(dateParam)
		// TODO add error condition check
		stopsFound := FilterByDate(dataFound.Data, uint32(date))

		context.JSON(http.StatusOK, stopsFound)
 	}	
}

func RangeDatesRoute(context *gin.Context) {
	area := context.Param("area")

	dataFound := selectData(area)

	result := rangeDataModel{
		DataBounds: dataFound.Bounds,
		StartDate: dataFound.StartDate,
		EndDate: dataFound.EndDate,
	}
	context.JSON(http.StatusOK, result)
	helpers.PrintMemresultUsage()
}

var GtfsInputData ConfigModel
func init() {
    // loads data
    GtfsInputData = ParseConfig()

}