package gtfstops

import (
	"encoding/json"
	"errors"
	"fmt"
	"gtfs_viewer/src/internals/utils"
	"io"
	"os"
	"runtime"
	"time"
)


type Stop struct {
	RouteLongName	string    	`json:"route_long_name"`
	Xcoord        	float32    	`json:"x"`
	Ycoord		 	float32    	`json:"y"`
	StartDate  		uint32    	`json:"start_date"`
	EndDate     	uint32    	`json:"end_date"`
	RouteType  		uint8		`json:"route_type"`
}


func (s *Stop) IsDateValid(date uint32) bool {
	return s.StartDate <= date && s.EndDate >= date
}


func ReadJson(path string) []Stop {
	defer utils.TimeTrack(time.Now(), "readJson")

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic(path + " not found")
	}

	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	defer runtime.GC()

	// read jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// init features array
	var features []Stop

	// unmarshal the byteArray containing the jsonFile's content into 'features' defined above
	json.Unmarshal(byteValue, &features)

	return features
}


func FilterByDate(features []Stop, date uint32) []Stop {
	defer utils.TimeTrack(time.Now(), "filterByDate")

    var featuresFiltered []Stop
    for _, stop := range features {

        if stop.IsDateValid(date) {
            featuresFiltered = append(featuresFiltered, stop)
        }
    }
	return featuresFiltered
}