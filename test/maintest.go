package test

import (
	"courseonline/config"
	"courseonline/controller"
	"courseonline/server"
	"courseonline/services"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestHttpServer(t *testing.T, store services.Store) *server.HttpServer {
	config := config.LoadConfig(getConfigFileName(), "../")
	handlerCtrl := controller.NewControllerManager(store)
	router := server.CreateRouter(handlerCtrl, "dev")
	server := server.NewHttpServer(&config, store, router)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env != "" {
		return "northwind-" + env
	}

	return "northwind"
}
