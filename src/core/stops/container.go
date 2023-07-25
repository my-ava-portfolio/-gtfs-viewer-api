package stops

type RangeDataModel struct {
	StartDate, EndDate uint32
	DataBounds         [4]float32
}

type StopsContainer struct {
	Files []StopsFeature
}

func (cm *StopsContainer) GetAreas() []string {
	filesCount := len(cm.Files)
	availableAreas := make([]string, filesCount)
	for index, feature := range cm.Files {
		availableAreas[index] = feature.Title
	}
	return availableAreas
}

func (cm StopsContainer) GetAreaRouteTypes(area string) []uint8 {
	dataFound := cm.selectData(area)
	return dataFound.RouteTypes
}

func (cm StopsContainer) GetRangesData(area string) RangeDataModel {
	dataFound := cm.selectData(area)
	return RangeDataModel{
		DataBounds: dataFound.Bounds,
		StartDate:  dataFound.StartDate,
		EndDate:    dataFound.EndDate,
	}
}

func (cm StopsContainer) GetStopsFilteredData(area string, date uint32, bounds []float32) []StopItemFiltered {
	dataFound := cm.selectData(area)

	var featuresFiltered []StopItemFiltered
	for _, stop := range dataFound.Data {

		if stop.IsDateValid(date) && stop.IntersectsBounds(bounds) {
			outputStop := StopItemFiltered{
				X: stop.Xcoord,
				Y: stop.Ycoord,
				RouteType: stop.RouteType,
				RouteId: stop.RouteId,
			}
			featuresFiltered = append(featuresFiltered, outputStop)
		}
	}
	return featuresFiltered
}

func (cm StopsContainer) selectData(area string) StopsFeature {
	var dataFound StopsFeature

	for _, feature := range cm.Files {
		if feature.Title == area {
			dataFound = feature
			break
		}
	}
	return dataFound
}