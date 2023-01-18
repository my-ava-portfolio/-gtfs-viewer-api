package stops

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"gtfs_viewer/src/helpers"
	"gtfs_viewer/src/internals/bounds"
	"gtfs_viewer/src/internals/uniques"
)

func GetData(path string, suffixFilter string) StopsContainer {
	var filesFound []StopsFeature

	files, err := os.ReadDir(path)
	if err != nil {
		panic(path + " directory not found")
	}
	for _, file := range files {
		if strings.Contains(file.Name(), suffixFilter) {
			fileSplit := strings.Split(file.Name(), suffixFilter)
			dataFound := readJson(path + file.Name())

			fileItem := StopsFeature{
				Title: fileSplit[0],
				Data:  dataFound,
			}
			computeDataMetadata(&fileItem)

			filesFound = append(filesFound, fileItem)
		}
	}
	config := StopsContainer{
		Files: filesFound,
	}
	return config
}

func computeDataMetadata(fileMetadata *StopsFeature) {
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
	fileMetadata.RouteTypes = uniques.Uint8(routeTypes)
}

func readJson(path string) []StopItem {
	defer helpers.TimeTrack(time.Now(), "ReadStopJson from "+path)

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
	var rawfeatures []StopItem

	decoder.Decode(&rawfeatures)
	features := make([]StopItem, len(rawfeatures))

	copy(features, rawfeatures)

	return features
}



