package main

import (
	timeHelper "gtfs_viewer/src/internals/time"
	"gtfs_viewer/src/structures"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)



func readJson(path string) ([]structures.Feature) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("not found")
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	// read jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// init features array
	var features []structures.Stop

	// unmarshal the byteArray containing the jsonFile's content into 'features' defined above
	json.Unmarshal(byteValue, &features)

	// set interface to each stop
	y := make([]structures.Feature, len(features))
	for i, v := range features {
		y[i] = v
	}

	return y
}

func filterByDate(features []structures.Feature, date int32) []structures.Feature {
	defer timeHelper.TimeTrack(time.Now(), "filterByDate")

    var featuresFiltered []structures.Feature
    for _, stop := range features {

        if stop.IsDateValid(date) {
            featuresFiltered = append(featuresFiltered, stop)
        }
    }
	return featuresFiltered
}


func main() {

	data := readJson("ter_data.json")
	dataFound := filterByDate(data, 1637856404)
	log.Println(strconv.Itoa(len(dataFound)))

}