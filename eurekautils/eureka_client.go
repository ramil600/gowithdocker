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

var heartBeatString = "http://discovery:8761/eureka/apps/dispatcher/"
var registerServiceString = "http://discovery:8761/eureka/apps/dispatcher" 




func (i InstanceInfo )RegisterService() {

	
	client := http.Client{}
	ticket := RegistrationTicket{Instance: i}

	ticketBody, err := json.Marshal(ticket)
	fmt.Println("ticket body:")
	fmt.Println(string(ticketBody))
	regbody := bytes.NewReader(ticketBody)
	if err != nil {
		log.Fatal("Cannot unmarshal instance")
	}
	fmt.Println("Instance :", ticket.Instance)
	fmt.Println("Ip Address:", ticket.Instance.HostName)
	req, err := http.NewRequest("POST", registerServiceString, regbody)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal("Could not form request")
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("could not send request")
	}

	fmt.Println("Response Status from eureka: ", res.Status)
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

	ins.Instance.HostName = getmyIPstr()
	ins.Instance.IpAddr = ins.Instance.HostName
	ins.Instance.HomePageUrl = "http://" + ins.Instance.IpAddr + ":8080"
	ins.Instance.StatusPageUrl = ins.Instance.HomePageUrl + "/status"
	ins.Instance.HealthCheckUrl = ins.Instance.HomePageUrl +"/healthcheck"

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
	
	sendString := heartBeatString + i.InstanceID

	fmt.Println("Sending heartbeat string: ..." )
	fmt.Println(sendString)

	req, err := http.NewRequest("PUT", sendString, nil)

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
