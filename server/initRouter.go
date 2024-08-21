package server

import (
	"courseonline/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRouter(handlers *controller.ControllerManager, mode string) *gin.Engine {
	var router *gin.Engine
	if mode == "test" {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	} else {
		router = gin.Default()
	}

	//router := gin.Default()
	//set a lower memory limit for multipart forms
	router.MaxMultipartMemory = 8 << 20 //8 Mib
	router.Static("/static", "./public")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))

	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Course Online")
	})

	return router

}
