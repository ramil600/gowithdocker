package eurekautils

type EurekaService interface {
	RegisterService()
	SendHeartBeat()
	ShutDown()
}