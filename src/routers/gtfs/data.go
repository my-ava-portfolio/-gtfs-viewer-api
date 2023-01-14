package gtfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"

	"gtfs_viewer/src/helpers"
	"gtfs_viewer/src/internals"
)

func getData(path string, suffixFilter string) []FileModel {
	var filesFound []FileModel

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic("Read directory issue")
	}
	for _, file := range files {
		if strings.Contains(file.Name(), suffixFilter) {
			fileSplit := strings.Split(file.Name(), suffixFilter)
			dataFound := readJson(path + file.Name())

			var x, y []float32
			var Dates []uint32
			for _, feature := range dataFound {
				x = append(x, feature.Xcoord)
				y = append(y, feature.Ycoord)
				Dates = append(Dates, feature.StartDate, feature.EndDate)
			}
			DatesBounds := internals.GetMinmax_uint32Array(Dates)
			bounds := helpers.GetBoundsFromXsAndYs(x, y)

			fileItem := FileModel{
				Title:     fileSplit[0],
				Data:      dataFound,
				Bounds:    bounds,
				StartDate: DatesBounds.Min,
				EndDate:   DatesBounds.Max}
			filesFound = append(filesFound, fileItem)
		}
	}

	return filesFound
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

