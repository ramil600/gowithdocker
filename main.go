package main

import (
	
	"time"
	"github.com/ramil600/gowithdocker/serviceone/eurekautils"
)

func main() {	
	eurekautils.RegisterService()
	time.Sleep(10 * time.Second)
	eurekautils.SendHeartBeat()
	time.Sleep(5 * time.Second)
	eurekautils.ShutDown()
}
