package gtfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"gtfs_viewer/src/helpers"
	"gtfs_viewer/src/internals/uniques"
	"gtfs_viewer/src/internals/bounds"

)

func GetData(path string, suffixFilter string) ConfigModel {
	var filesFound []FileModel

	files, err := os.ReadDir(path)
	if err != nil {
		panic(path + " directory not found")
	}
	for _, file := range files {
		if strings.Contains(file.Name(), suffixFilter) {
			fileSplit := strings.Split(file.Name(), suffixFilter)
			dataFound := readJson(path + file.Name())

			fileItem := FileModel{
				Title:     fileSplit[0],
				Data:      dataFound,
			}
			computeDataMetadata(&fileItem)

			filesFound = append(filesFound, fileItem)
		}
	}
	config := ConfigModel{
		Files: filesFound,
	}
	return config
}

func computeDataMetadata(fileMetadata *FileModel) {
	var x, y []float32
	var dates []uint32
	var routeTypes []uint8
	for _, feature := range fileMetadata.Data {
		x = append(x, feature.Xcoord)
		y = append(y, feature.Ycoord)
		dates = append(dates, feature.StartDate, feature.EndDate)
		routeTypes = append(routeTypes, feature.RouteType)
	}

	DatesBounds := bounds.GetMinmax_uint32Array(dates)
	fileMetadata.StartDate = DatesBounds.Min
	fileMetadata.EndDate = DatesBounds.Max
	fileMetadata.Bounds = helpers.GetBoundsFromXsAndYs(x, y)
	fileMetadata.routeTypes = uniques.Uint8(routeTypes)
}

func readJson(path string) []Stop {
	defer helpers.TimeTrack(time.Now(), "ReadStopJson from " + path)

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
	decoder := json.NewDecoder(jsonFile)
	
	// init features array
	var rawfeatures []Stop

    decoder.Decode(&rawfeatures)
	features := make([]Stop, len(rawfeatures))

	copy(features, rawfeatures)

	return features
}

func SelectData(area string) FileModel {
	var dataFound FileModel

	for _, feature := range gtfsInputData.Files {
		if feature.Title == area {
			dataFound = feature
			break
		}
	}
	return dataFound
}

func (s *Stop) IsDateValid(date uint32) bool {
	return s.StartDate <= date && s.EndDate >= date
}

func FilterByDate(features []Stop, date uint32, bounds [4]float32) []Stop {
	defer helpers.TimeTrack(time.Now(), "filterByDate")

    var featuresFiltered []Stop
    for _, stop := range features {

        if stop.IsDateValid(date) && helpers.IsPointIntersectsBounds([2]float32{stop.Xcoord, stop.Ycoord}, bounds) {
            featuresFiltered = append(featuresFiltered, stop)
        }
    }
	return featuresFiltered
}


var gtfsInputData ConfigModel
func init() {
    // loads data
	//gtfsInputData = ParseConfig(&ModeSelected)

    //GtfsInputData = ParseConfig()

}
