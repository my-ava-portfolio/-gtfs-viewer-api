package main

import (

	gtfs "gtfs_viewer/src/routers/gtfs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


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
	  
	gtfs.GtfsGroupRouterRequests(router)

	router.Run(":7001")
}