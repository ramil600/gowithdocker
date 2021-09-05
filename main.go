package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ramil600/gowithdocker/serviceone/eurekautils"
	"github.com/ramil600/gowithdocker/serviceone/handlers"
)

func main() {	
	ins := eurekautils.GetNewInstance()
	ins.RegisterService()
	time.Sleep(10 * time.Second)
	
	e := echo.New()
	e.GET("/ping", handlers.PingHandler)
	e.Logger.Fatal(e.Start(":8000"))
	
	for i:=0; i < 2; i++ {
		ins.SendHeartBeat()
		time.Sleep(5 * time.Second)
	}

	ins.ShutDown()
	
}
