package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	gtfs "gtfs_viewer/src/routers/gtfs"
)

var Router *gin.Engine
func init() {
	Router = setupRouter()
	gtfs.GtfsGroupRouterRequests(Router)

}

func TestMovingNodesRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/gtfs_builder/ter/moving_nodes?date=1637857000", nil)
	Router.ServeHTTP(w, req)

    var stops []gtfs.Stop
    json.Unmarshal(w.Body.Bytes(), &stops)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, stops)
}

func TestRangeDatesRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/gtfs_builder/ter/range_dates", nil)
	Router.ServeHTTP(w, req)

    var rangeData gtfs.RangeDataModel
    json.Unmarshal(w.Body.Bytes(), &rangeData)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, rangeData)
}

func TestRouteTypesRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/gtfs_builder/ter/route_types", nil)
	Router.ServeHTTP(w, req)

    var routeTypes []uint8
    json.Unmarshal(w.Body.Bytes(), &routeTypes)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, routeTypes)
}

func TestAvailableAreasRoute(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v2/gtfs_builder/existing_study_areas", nil)
	Router.ServeHTTP(w, req)

    var availableAreas []string
    json.Unmarshal(w.Body.Bytes(), &availableAreas)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, availableAreas)
}
