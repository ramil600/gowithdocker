package eurekautils

type Port struct {
	Port    string  `json:"$"`
	Enabled string `json:"@enabled"`
}

type DataCenterInfo struct {
	Name  string `json:"name"`
	Class string `json:"@class"`
}

type InstanceInfo struct {
	HostName                      string          `json:"hostName"`
	HomePageUrl                   string          `json:"homePageUrl,omitempty"`
	StatusPageUrl                 string          `json:"statusPageUrl"`
	HealthCheckUrl                string          `json:"healthCheckUrl,omitempty"`
	App                           string          `json:"app"`
	IpAddr                        string          `json:"ipAddr"`
	VipAddress                    string          `json:"vipAddress"`
	SecureVipAddress              string          `json:"secureVipAddress,omitempty"`
	Status                        string          `json:"status"`
	Port                          *Port           `json:"port,omitempty"`
	SecurePort                    *Port           `json:"securePort,omitempty"`
	DataCenterInfo                *DataCenterInfo `json:"dataCenterInfo"`
	IsCoordinatingDiscoveryServer bool            `json:"isCoordinatingDiscoveryServer,omitempty"`
	LastUpdatedTimestamp          int             `json:"lastUpdatedTimestamp,omitempty"`
	LastDirtyTimestamp            int             `json:"lastDirtyTimestamp,omitempty"`
	ActionType                    string          `json:"actionType,omitempty"`
	Overriddenstatus              string          `json:"overriddenstatus,omitempty"`
	CountryId                     int             `json:"countryId,omitempty"`
	InstanceID                    string          `json:"instanceId,omitempty"`
}

type RegistrationTicket struct {
	Instance InstanceInfo `json:"instance"`
}
