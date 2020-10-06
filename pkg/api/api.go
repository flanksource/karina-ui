package api

import (
	"time"
)

type Cluster struct {
	Name         string     `json:"name,omitempty"`
	Properties   Properties `json:"properties,omitempty"`
	CanaryChecks []Chek     `json:"canary_checks,omitempty"`
	Nodes        []Node     `json:"nodes,omitempty"`
	Alerts       []Alert    `json:"alerts,omitempty"`
}

type Properties struct {
	Capacity 		Capacity		`json:"capacity,omitempty"`
	Allocatable 	Allocatable 	`json:"allocatable,omitempty"`
	Condititions 	[]Condition 	`json:"conditions,omitempty"`
	Addresses 		[]Address 		`json:"addresses,omitempty"`
	DaemonEndpoints KubeletEndpoint `json:"daemonendpoints,omitempty"`
	NodeInfo 		Nold 			`json:"nodeInfo,omitempty"`
	Alerts			[]Alert 		`json:"alerts,omitempty"`
}

type Capacity struct {
	Cpu 	string `json:"cpu,omitempty"`
	EphStor string `json:"ephemeral-storage,omitempty"`
	OneGi 	string `json:"hugepages-1Gi,omitempty"`
	TwoMi 	string `json:"hugepages-2Mi,omitempty"`
	Memory 	string `json:"memory,omitempty"`
	Pods 	string `json:"pods,omitempty"`
}

type Allocatable struct {
	Cpu 	string `json:"cpu,omitempty"`
	EphStor string `json:"ephemeral-storage,omitempty"`
	OneGi 	string `json:"hugepages-1Gi,omitempty"`
	TwoMi 	string `json:"hugepages-2Mi,omitempty"`
	Memory 	string `json:"memory,omitempty"`
	Pods 	string `json:"pods,omitempty"`
}

type Condition struct {
	Type 			string 		`json:"type,omitempty"`
	Status 			string 		`json:"status,omitempty"`
	LastHeartbeat 	time.Time 	`json:"lastHeartbeatTime,omitempty"`
	LastTransition 	time.Time 	`json:"lastTransitionTime,omitempty"`
	Reason 			string 		`json:"reason,omitempty"`
	Message 		string 		`json:"message,omitempty"`
}


type Address struct {
	Type 	string `json:"type,omitempty"`
	Address string `json:"type,omitempty"`
}


type KubeletEndpoint struct {
	Port string `json:"Port,omitempty"`
}


type Canarydata struct {
	ServerName string `json:"server_name,omitempty"`
	Checks     []Chek `json:"checks,omitempty"`
}

type Chek struct {
	Key           string       `json:"key,omitempty"`
	Type          string       `json:"type,omitempty"`
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	Uptime        string       `json:"uptime,omitempty"`
	Latency       string       `json:"latency,omitempty"`
	CheckStatuses []ChekStatus `json:"checkStatuses,omitempty"`
	checkConf     ChekConf     //`json:"checkConf,omitempty"`

}

type ChekStatus struct {
	Status   bool
	Invalid  bool
	Time     string
	Duration int
	Message  string
}

type ChekConf struct {
	Server          string
	Port            int
	Query           string
	QueryType       string
	Minrecords      string
	ExactReply      []string
	Timeout         int
	ThresholdMillis int
}


type Nold struct {
	MachineID 				string `json:"machineID,omitempty"`
	SystemUUID 				string `json:"systemUUID,omitempty"`
	BootID 					string `json:"bootID,omitempty"`
	KernelVersion 			string `json:"kernelVersion,omitempty"`
	OsImage 				string `json:"osImage,omitempty"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion,omitempty"`
	KubeletVersion 			string `json:"kubeletVersion,omitempty"`
	KubeProxyVersion 		string `json:"kubeProxyVersion,omitempty"`
	OperatingSystem 		string `json:"operatingSystem,omitempty"`
	Architecture 			string `json:"architecture,omitempty"`
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
	Type 	string    `json:"type,omitempty"`
	Cluster string    `json:"cluster,omitempty"`
}

type ClusterConfiguration struct {
	CanaryChecker string `yaml:"canaryChecker,omitempty"`
	Prometheus    string `yaml:"prometheus,omitempty"`
	AlertManager  string `yaml:"alertmanager,omitempty"`
	Kubeconfig    string `yaml:"kubeconfig,omitempty"`
}

/*type Combo struct {
	Cluster []Cluster
	Canary 	Canarydata
}*/

/*
type Property struct {
	Name   string  `json:"name,omitempty"`
	Value  string  `json:"label,omitempty"`
	Type   string  `json:"type,omitempty"`
	Icon   string  `json:"icon,omitempty"`
	Alerts []Alert `json:"alerts,omitempty"`
}

type CanaryCheck struct {
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
