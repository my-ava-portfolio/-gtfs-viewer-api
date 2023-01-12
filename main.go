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



func readJson(path string) ([]structures.Stop) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("not found")
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

	// read jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// init features array
	var features []structures.Stop

	// unmarshal the byteArray containing the jsonFile's content into 'features' defined above
	json.Unmarshal(byteValue, &features)

	// we iterate through every feature
	//for i := 0; i < len(features); i++ {
		//log.Printf("x: " + strconv.FormatFloat(features[i].x, 'g', 3, 64) + "- y: " + strconv.FormatFloat(features[i].y, 'g', 3, 64))
	//}
	return features

}

func filterByDate(features []structures.Stop, date int32) []structures.Stop {
	defer timeHelper.TimeTrack(time.Now(), "filterByDate")

    var featuresFiltered []structures.Stop

    for _, stop := range features {

        if stop.StartDate <= date && stop.EndDate >= date {
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