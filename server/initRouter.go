package server

import (
	"courseonline/controller"
	"courseonline/middleware"

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

	categoryRoute := api.Group("/category")
	{
		//categoryRoute.Use(middleware.AuthMiddleware())
		categoryRoute.GET("/", handlers.GetListCategory)
		categoryRoute.POST("/", handlers.PostCategory)
		categoryRoute.PUT("/:id", handlers.UpdateCategory)
		categoryRoute.GET("/:id", handlers.GetCategoryById)

	}

	courseRoute := api.Group("/course")
	{
		//courseRoute.Use(middleware.AuthMiddleware())
		courseRoute.GET("/", handlers.GetListCourse)
		courseRoute.POST("/", handlers.PostCourse)
		courseRoute.GET("/:id", handlers.GetCourseById)
		courseRoute.PUT("/:id", handlers.UpdateCourse)
		courseRoute.DELETE("/:id", handlers.DeleteCourse)
	}

	courseimageRoute := api.Group("/image")
	{
		courseimageRoute.Use(middleware.AuthMiddleware())
		courseimageRoute.GET("/", handlers.GetListImage)
		courseimageRoute.POST("/", handlers.CreateCourseimages)
		courseimageRoute.GET("/:id", handlers.FindCourseImageByID)
		courseimageRoute.PUT("/:id", handlers.UpdateCourseImage)
		courseimageRoute.DELETE("/:id", handlers.DeleteCourseImage)
	}

	userRoute := api.Group("/user")
	{
		userRoute.POST("/signup", handlers.Signup)
		userRoute.POST("/signin", handlers.Sigin)
		userRoute.POST("/signout", handlers.Signout)
		userRoute.GET("/profile", handlers.GetProfile)
	}

	cartRoute := api.Group("/cart")
	{
		cartRoute.POST("/", handlers.AddToCart)
		cartRoute.GET("/", handlers.FindAllCart)
		cartRoute.GET("/:id", handlers.FindCartById)
		cartRoute.POST("/order", handlers.CreateCart)
		cartRoute.GET("/user:id", handlers.FindCartByCartUserID)
	}

	ordercoursedetailRoute := api.Group("/detail")
	{
		ordercoursedetailRoute.GET("/", handlers.GetListOrderCourseDetail)
		ordercoursedetailRoute.POST("/", handlers.PostOrderCourseDetail)
		ordercoursedetailRoute.PUT("/:id", handlers.UpdateOrderCourseDetail)
		ordercoursedetailRoute.DELETE("/:id", handlers.DeleteOrderCourseDetail)
	}

	return router

}
