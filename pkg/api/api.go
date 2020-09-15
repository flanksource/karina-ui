package api

import (
	"time"
	
)

type Cluster struct {
	Name         string      `json:"name,omitempty"`
	Properties   []Property  `json:"properties,omitempty"`
	CanaryChecks []CanaryCheck `json:"canary_checks,omitem/pty"`
	Nodes        []Node      `json:"nodes,omitempty"`
	Alerts       []Alert     `json:"alerts,omitempty"`
}

type Property struct {
	Name   string  `json:"name,omitempty"`
	Value  string  `json:"value,omitempty"`
	Type   string  `json:"type,omitempty"`
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
