package config

import (
	"gtfs_viewer/src/gtfstops"
	"os"
	"strings"
)


type FileModel struct {
	Title		string
	Data		[]gtfstops.Stop
}

type ConfigModel struct {
	Files 		[]FileModel
}



func ParseConfig() ConfigModel {
	filesFound := getFiles("data/", "_data.json")

	config := ConfigModel{
		Files: 		filesFound,
	}
	return config
}

func getFiles(path string, suffixFilter string) []FileModel {
	var filesFound []FileModel

	files, err := oSReadDir(path)
    if err != nil {
        panic(err)
    }
	for _, file := range files {
		if strings.Contains(file, suffixFilter) {
			fileSplit := strings.Split(file, suffixFilter)
			fileItem := FileModel{
				Title: fileSplit[0], 
				Data: gtfstops.ReadJson(path + file),
			}
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