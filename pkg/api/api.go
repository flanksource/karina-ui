package api

import (
	"time"
)

type Cluster struct {
	Name         string      `json:"name,omitempty"`
	Properties   []Property  `json:"properties,omitempty"`
	CanaryChecks []Chek  	 `json:"canary_checks,omitempty"`
	Nodes        []Node      `json:"nodes,omitempty"`
	Alerts       []Alert     `json:"alerts,omitempty"`
}

type Property struct {
	Name   string  `json:"name,omitempty"`
	Value  string  `json:"value,omitempty"`
	Type   string  `json:"type,omitempty"`
	Alerts []Alert `json:"alerts,omitempty"`
}


type Canarydata struct {
	ServerName	string 	`json:"server_name,omitempty"`
	Checks 		[]Chek 	`json:"checks,omitempty"`
}

type Chek struct {
	Key 			string 			`json:"key,omitempty"`
	Type 			string 			`json:"type,omitempty"`
	Name  			string 			`json:"name,omitempty"`
	Description 	string 			`json:"description,omitempty"`
	Uptime 			string 			`json:"uptime,omitempty"`
	Latency 		string 			`json:"latency,omitempty"`
	CheckStatuses 	[]ChekStatus 	`json:"checkStatuses,omitempty"`
	checkConf 		ChekConf 		//`json:"checkConf,omitempty"`

}

type ChekStatus struct {

	Status 		bool
	Invalid 	bool
	Time  		string
	Duration	int
	Message 	string
}

type ChekConf struct {
	Server 			string
	Port 			int
	Query 			string
	QueryType 		string
	Minrecords 		string
	ExactReply 		[]string
	Timeout   		int
	ThresholdMillis int

}

type Node struct {
	Name   string  `json:"name,omitempty"`
	IP     string  `json:"ip,omitempty"`
	Alerts []Alert `json:"alerts,omitempty"`
}

type Alert struct {
	Level   string    `json:"level,omitempty"`
	Since   time.Time `json:"since,omitempty"`
	Message string    `json:"message,omitempty"`
}

type ClusterConfiguration struct {
	CanaryChecker string `yaml:"canaryChecker,omitempty"`
	Prometheus    string `yaml:"prometheus,omitempty"`
	AlertManager  string `yaml:"alertmanager,omitempty"`
	Kubeconfig 	  string `yaml:"kubeconfig,omitempty"`
}

/*type Combo struct {
	Cluster []Cluster
	Canary 	Canarydata
}*/

/*type CanaryCheck struct {
	Name      string              `json:"name,omitempty"`
	Namespace string              `json:"namespace,omitempty"`
	Type      string              `json:"type,omitempty"`
	Endpoint  string              `json:"endpoint,omitempty"`
	Latency1H string              `json:"latency_1_h,omitempty"`
	Uptime1H  string              `json:"uptime_1_h,omitempty"`
	Status    []CanaryCheckStatus `json:"status,omitempty"`
}

type CanaryCheckStatus struct {
	Pass      bool      `json:"pass,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Message   string    `json:"message,omitempty"`
}*/