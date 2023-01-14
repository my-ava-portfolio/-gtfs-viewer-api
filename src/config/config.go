package config

import (
	"gtfs_viewer/src/gtfstops"
	"gtfs_viewer/src/helpers"
	"gtfs_viewer/src/internals"
	"os"
	"strings"
)

type FileModel struct {
	Title     string
	Data      []gtfstops.Stop
	Bounds    [4]float32
	StartDate uint32
	EndDate   uint32
}

type ConfigModel struct {
	Files []FileModel
}

func ParseConfig() ConfigModel {
	filesFound := getData("data/", "_data.json")

	config := ConfigModel{
		Files: filesFound,
	}
	return config
}

func getData(path string, suffixFilter string) []FileModel {
	var filesFound []FileModel

	files, err := oSReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file, suffixFilter) {
			fileSplit := strings.Split(file, suffixFilter)
			dataFound := gtfstops.ReadJson(path + file)

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

func oSReadDir(path string) ([]string, error) {
	var files []string
	f, err := os.Open(path)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
