package main

import (
	"gtfs_viewer/src/routers/gtfs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	"github.com/mvrilo/go-redoc"
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

	doc := redoc.Redoc{
		Title:       "gtfs-viewer-api",
		Description: "",
		SpecFile:    "./openapi.json", // "./openapi.yaml"
		SpecPath:    "/openapi.json",  // "/openapi.yaml"
		DocsPath:    "/schema",
	}

	gin.SetMode(gin.ReleaseMode)    

	router := setupRouter()

	gtfs.GtfsGroupRouterHandler("data/", router)
	router.Use(ginredoc.New(doc))

	router.Run(":7001")
}