package main

import (
	"fmt"
	snow "github.com/FracKenA/gosnow"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


var APIPort = 3808
var LogLevel = "DEBUG"
var MD5Seed = "BlahBlah"
var SNOWInstance = "dev87192"
var SNOWUsername = "admin"
var SNOWPassword = "x25EWaGMcdGy"
var SNOWDomain = "service-now.com"
var SNOWConnect = snow.InitializeConnection(SNOWInstance,SNOWDomain,SNOWUsername,SNOWPassword)
var SNOWImpact = 1
var SNOWUrgency = 1
var SNOWAssignmentGroup = "Help Desk"
var SNOWCategory = "Phone"
var SNOWCaller = "admin"

func main() {
	// API interface creation
	echoInterface := echo.New()

	// Establish Middleware functions
	echoInterface.Use(middleware.Recover())
	echoInterface.Use(middleware.Logger())

	// Hide Echo Framework banner
	echoInterface.HidePort = true
	echoInterface.HideBanner = true

	// Configure endpoint versioning
	aPIEndPoint := echoInterface.Group("/api")
	v1Route := aPIEndPoint.Group("/v1")

	// Configure v1 endpoints
	v1Route.POST("/incident", PostV1Incident)

	// Start the API Server
	binding := fmt.Sprintf(":%d",APIPort)

	echoInterface.Logger.Fatal(echoInterface.Start(binding))

}
