package api

import (
	"time"
)

type Cluster struct {
	Name         string     		`json:"name,omitempty"`
	Properties   []Property 		`json:"properties,omitempty"`
	CanaryChecks []Chek     		`json:"canary_checks,omitempty"`
	Nodes        []Node    			`json:"nodes,omitempty"`
	Alerts       []Alerti  `json:"alerts,omitempty"`
	Prometheus   []PrometheusData 	`json:"prometheus,omitempty"`

}

type Property struct {
	Name   string  `json:"name,omitempty"`
	Value  string  `json:"label,omitempty"`
	Type   string  `json:"type,omitempty"`
	Icon   string  `json:"icon,omitempty"`
	Alerts []Alert `json:"alerts,omitempty"`
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
	Kubeconfig    string `yaml:"kubeconfig,omitempty"`
}


type PrometheusAlert struct {
	Status 	string 		`json:"status,omitempty"`
	Data 	AlertData		`json:"data,omitempty"`
}

type AlertData struct {
	Alerts 	[]Alerti
}

type Alerti struct {
	Label 		Label		`json:"labels,omitempty"` 
	Annotation Annotation 	`json:"annotations,omitempty"`
	Since 	time.Time 		`json:"since,omitempty"`
	//Value 		[]interface{}
}

type Annotation struct {
	Message 	string 	`json:"message,omitempty"`
}


type Label struct {
	AlertName 	string 	`json:"alertname,omitempty"`
	Level 		string 	`json:"severity,omitempty"`
}

type PrometheusData struct {
	//Status 	string 		`json:"status,omitempty"`
	Data 	Data		`json:"data,omitempty"`
}

type Data struct {
	//ResultType 	string 	`json:"resultType,omitempty"`
	Result 		[]Result 	`json:"result,omitempty"`
}


type Result struct {
	Metric Metric 			`json:"metric,omitempty"`
	Value  []interface{} 	`json:"value,omitempty"`
}

type Metric struct {
	Name 		string `json:"__name__,omitempty"`
	Endpoint 	string `json:"endpoint,omitempty"`
	Instance 	string `json:"instance,omitempty"`
	Job 		string `json:"job,omitempty"`
	MetricsPath string `json:"metrics_path,omitempty"`
	Namespace 	string `json:"namespace,omitempty"`
	Node 		string `json:"node,omitempty"`
	Service 	string `json:"service,omitempty"`
	Pod 		string `json:"pod,omitempty"`
	Device 		string `json:"device,omitempty"`
	FSType 		string `json:"fstype,omitempty"`
	MountPoint 	string `json:"mountpoint,omitempty"`
}