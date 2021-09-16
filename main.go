package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ramil600/gowithdocker/serviceone/eurekautils"
	"github.com/ramil600/gowithdocker/serviceone/handlers"
)

func main() {	
	ins := eurekautils.GetNewInstance()
	ins.RegisterService()
	time.Sleep(50 * time.Second)

	e := echo.New()
	e.GET("/ping", handlers.PingHandler)
	e.GET("/", handlers.IndexHandler)
	e.GET("/temperature/:uuid/:temp", handlers.TempHandler)

	f, ferr := os.Create("logger.txt")

	if ferr != nil {
		 log.Fatal(ferr)
	}
	defer f.Close()

	go func() {
		for  {
			ins.SendHeartBeat()
			time.Sleep(50 * time.Second)
		}

	}()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method} status=${status} uri=${uri}\n", 
		Output: f,
	}))
	e.Logger.Fatal(e.Start(":8080"))
	
	

	//ins.ShutDown()
	
}
