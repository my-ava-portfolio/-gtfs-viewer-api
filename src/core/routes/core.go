package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
	"gtfs_viewer/src/helpers"
)

func GetData(path string, suffixFilter string) RoutesContainer {
	var filesFound []RoutesFeature

	files, err := os.ReadDir(path)
	if err != nil {
		panic(path + " directory not found")
	}
	for _, file := range files {
		if strings.Contains(file.Name(), suffixFilter) {
			fileSplit := strings.Split(file.Name(), suffixFilter)
			dataFound := readJson(path + file.Name())

			fileItem := RoutesFeature{
				Title: fileSplit[0],
				Data:  dataFound,
			}
			filesFound = append(filesFound, fileItem)
		}
	}
	config := RoutesContainer{
		Files: filesFound,
	}
	return config
}

func readJson(path string) []StopRouteItem {
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
	var rawfeatures []StopRouteItem

	decoder.Decode(&rawfeatures)
	features := make([]StopRouteItem, len(rawfeatures))

	copy(features, rawfeatures)

	return features
}



