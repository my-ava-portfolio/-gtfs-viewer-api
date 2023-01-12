package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)


func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

type Stop struct {
	Index  			string    	`json:"index"`
	Xcoord        	float64    	`json:"x"`
	Ycoord		 	float64    	`json:"y"`
	StartDate  		int32    	`json:"start_date"`
	EndDate     	int32    	`json:"end_date"`
	RouteType  		string		`json:"route_type"`
	RouteLongName	string    	`json:"route_long_name"`
}


func readJson(path string) ([]Stop) {
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

    //var res map[string]interface{}
    //json.Unmarshal([]byte(byteValue), &res)

	// init features array
	var features []Stop

	// unmarshal the byteArray containing the jsonFile's content into 'features' defined above
	json.Unmarshal(byteValue, &features)

	// we iterate through every feature
	//for i := 0; i < len(features); i++ {
		//log.Printf("x: " + strconv.FormatFloat(features[i].x, 'g', 3, 64) + "- y: " + strconv.FormatFloat(features[i].y, 'g', 3, 64))
	//}
	return features

}

func filterByDate(features []Stop, date int32) []Stop {
	defer timeTrack(time.Now(), "filterByDate")

    var featuresFiltered []Stop

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