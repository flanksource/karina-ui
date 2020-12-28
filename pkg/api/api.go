package api

import (
	"time"
)

type Cluster struct {
	Name         string     `json:"name,omitempty"`
	Kubernetes   Kubernetes `json:"kubernetes,omitempty"`
	Properties   []Property `json:"properties,omitempty"`
	CanaryChecks []Check    `json:"canaryChecks,omitempty"`
	Alerts       []Alert    `json:"alerts,omitempty"`
}

type Property struct {
	Name   string  `json:"name,omitempty"`
	Value  string  `json:"label,omitempty"`
	Type   string  `json:"type,omitempty"`
	Icon   string  `json:"icon,omitempty"`
	Alerts []Alert `json:"alerts,omitempty"`
}

type Canarydata struct {
	ServerName string  `json:"server_name,omitempty"`
	Checks     []Check `json:"checks,omitempty"`
}

type Check struct {
	Key           string        `json:"key,omitempty"`
	Type          string        `json:"type,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Uptime        string        `json:"uptime,omitempty"`
	Latency       string        `json:"latency,omitempty"`
	CheckStatuses []CheckStatus `json:"checkStatuses,omitempty"`
}

type CheckStatus struct {
	Status   bool
	Invalid  bool
	Time     string
	Duration int
	Message  string
}

type Node struct {
	Name             string  `json:"name"`
	InternalIP       string  `json:"internalIP"`
	ExternalIP       string  `json:"externalIP"`
	KernelVersion    string  `json:"kernelVersion"`
	KubeletVersion   string  `json:"kubeletVersion"`
	ContainerRuntime string  `json:"containerRuntime"`
	OSImage          string  `json:"osImage"`
	Alerts           []Alert `json:"alerts"`
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

type Config struct {
	Clusters map[string]ClusterConfiguration `yaml:"clusters,omitempty"`
}

type Kubernetes struct {
	CPU            int     `json:"cpu"`
	Memory         int64   `json:"memory"`
	Disk           int64   `json:"disk"`
	KubeletVersion string  `json:"kubeletVersion"`
	KernelVersion  string  `json:"kernelVersion"`
	OSVersion      string  `json:"osVersion"`
	CRIVersion     string  `json:"criVersion"`
	Nodes          []Node  `json:"nodes"`
	KubeletAlerts  []Alert `json:"kubeletAlerts"`
	KernelAlerts   []Alert `json:"kernelAlerts"`
	OSAlerts       []Alert `json:"osAlerts"`
	CRIAlerts      []Alert `json:"criAlerts"`
}
