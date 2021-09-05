package main

import (
	"time"
	"github.com/ramil600/gowithdocker/serviceone/eurekautils"
)

func main() {	
	ins := eurekautils.GetNewInstance()
	ins.RegisterService()
	time.Sleep(10 * time.Second)
	ins.SendHeartBeat()
	time.Sleep(5 * time.Second)
	ins.ShutDown()
	
}
