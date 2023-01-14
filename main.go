package main

import (

	gtfs "gtfs_viewer/src/routers/gtfs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://portfolio.amaury-valorge.com"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
	  }))
	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)    

	router := setupRouter()
	  
	gtfs.GtfsGroupRouterRequests(router)

	router.Run(":7001")
}