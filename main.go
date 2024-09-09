package main

import (
	"courseonline/config"
	"courseonline/controller"
	"courseonline/server"
	"courseonline/services"
	"log"
	"os"
)

// @title           CourseOnline
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   tqrcmn30@gmail.com
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8888
// @BasePath  /api/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	log.Println("Starting Courseonline App")
	log.Println("Initializing configuration")

	config := config.LoadConfig(getConfigFileName(), ".")
	log.Println("Initializing database")
	dbConnection := server.InitDatabase(&config)
	defer server.Close(dbConnection)

	store := services.NewStoreManager(dbConnection)

	handlerCtrl := controller.NewControllerManager(store)

	router := server.CreateRouter(handlerCtrl, "dev")

	log.Println("Initializig HTTP sever")
	httpServer := server.NewHttpServer(&config, store, router)

	httpServer.MountSwaggerHandlers()
	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env != "" {
		return "northwind-" + env
	}

	return "northwind"
}
