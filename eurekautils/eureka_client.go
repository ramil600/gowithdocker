package eurekautils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var heartBeatString = "http://localhost:8761/eureka/apps/vendor3/WKS-SOF-L011"





func (i InstanceInfo )RegisterService() {

	regbody := bytes.NewReader(parseTPL())
	client := http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8761/eureka/apps/vendor3", regbody)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal("Could not form request")
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("could not send request")
	}

	fmt.Println(res.Body)
	fmt.Println ("Instance host name: ",i.HostName)
}

func getmyIPstr() string {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic("Could not get all the addreses")
	}

	for _, addr := range addrs {

		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}
	return ""
}

func GetNewInstance() *InstanceInfo {
	var ins = RegistrationTicket{}
	body := parseTPL()
	err := json.Unmarshal(body, &ins)

	if err != nil {
		log.Fatal("Could not read the Registration Ticket")
	}

	newInstance := new(InstanceInfo)
	newInstance = &ins.Instance
	return newInstance


}

func parseTPL() []byte {
	var myinstance RegistrationTicket
	regtpl, err := os.ReadFile("reg_vendor.json")
	
	fmt.Println(myinstance.Instance.InstanceID)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(regtpl, &myinstance)
	fmt.Println(myinstance.Instance.InstanceID)
	fmt.Println(myinstance.Instance.HostName)

	if err != nil{
		log.Fatal(err)
	}
	return regtpl

}


func (i InstanceInfo) SendHeartBeat() {

	// NOTE: %s/eureka/apps/SERVICENAME/192.168.1.49:SERVICENAME:9000
	//fmt.Sprintf("%s/eureka/apps/%s/%s", s.EurekaService, s.RegistrationTicket.Instance.App, s.RegistrationTicket.Instance.InstanceId)
	client := http.Client{}
	//myIp := getmyIPstr()

	fmt.Println(heartBeatString)

	req, err := http.NewRequest("PUT", heartBeatString, nil)

	if err != nil {
		log.Fatal("Could not form request")
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("could not send request")
	}

	if res.StatusCode == http.StatusNotFound {
		log.Fatal("The application is not found")
	}
	fmt.Println("The status is ", res.StatusCode)

}

func (i InstanceInfo) ShutDown() {

	client := http.Client{}
	req, err := http.NewRequest("DELETE", heartBeatString, nil)

	if err != nil {
		log.Fatal("Could not form request")
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("could not send request")
	}

	if res.StatusCode == http.StatusNotFound {
		log.Fatal("The application is not found")
	}
	fmt.Println("The status is ", res.StatusCode)
}
