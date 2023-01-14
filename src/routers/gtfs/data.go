package gtfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"gtfs_viewer/src/helpers"
	"gtfs_viewer/src/internals/uniques"
	"gtfs_viewer/src/internals/bounds"

)

func getData(path string, suffixFilter string) []FileModel {
	var filesFound []FileModel

	files, err := os.ReadDir(path)
	if err != nil {
		panic("Read directory issue")
	}
	for _, file := range files {
		if strings.Contains(file.Name(), suffixFilter) {
			fileSplit := strings.Split(file.Name(), suffixFilter)
			dataFound := readJson(path + file.Name())

			fileItem := FileModel{
				Title:     fileSplit[0],
				Data:      dataFound}
			computeDataMetadata(&fileItem)

			filesFound = append(filesFound, fileItem)
		}
	}

	return filesFound
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
	defer helpers.TimeTrack(time.Now(), "readJson")

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

func SelectData(area string) FileModel {
	var dataFound FileModel

	for _, feature := range GtfsInputData.Files {
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

func FilterByDate(features []Stop, date uint32) []Stop {
	defer helpers.TimeTrack(time.Now(), "filterByDate")

    var featuresFiltered []Stop
    for _, stop := range features {

        if stop.IsDateValid(date) {
            featuresFiltered = append(featuresFiltered, stop)
        }
    }
	return featuresFiltered
}

var GtfsInputData ConfigModel
func init() {
    // loads data
    GtfsInputData = ParseConfig()

}
