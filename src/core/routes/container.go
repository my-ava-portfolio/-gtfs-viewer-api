package routes

import (
	"fmt"
	"strconv"
)

type RoutesContainer struct {
	Files []RoutesFeature
}

func (cm RoutesContainer) GetRouteNameByRouteId(area string, id string) string {
	dataFound := cm.selectData(area)
	longName := "missing"

	for _, route := range dataFound.Data {
		fmt.Print(route)

		if route.RouteId == strconv.ParseUint(id) {
			longName = route.RouteLongName
			break
		}
	}
	return longName
}

func (cm RoutesContainer) selectData(area string) RoutesFeature {
	var dataFound RoutesFeature

	for _, feature := range cm.Files {
		if feature.Title == area {
			dataFound = feature
			break
		}
	}
	return dataFound
}
