package main

import (

	gtfsRoute "gtfs_viewer/src/routers/gtfs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func gtfsGroupRouterRequests(router *gin.Engine) {
	v2 := router.Group("/api/v2/gtfs_builder")

	v2.GET(":area/moving_nodes", gtfsRoute.MovingStopsRoute)
	v2.GET(":area/range_dates", gtfsRoute.RangeDatesRoute)
	//v2.GET("/route_types", movingStopsRoute)

}



func main() {
	// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/24.html
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://portfolio.amaury-valorge.com"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	  }))
	  
	gtfsGroupRouterRequests(router)

	router.Run(":7001")
}