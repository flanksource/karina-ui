package ui

import (
	"fmt"
	"strconv"

	"github.com/flanksource/karina-ui/pkg/api"
	"github.com/pkg/errors"
)

const gb = 1024 * 1024 * 1024

func Transform(clusters []api.Cluster) ([]Cluster, error) {
	response := make([]Cluster, len(clusters))

	for i, cluster := range clusters {
		c, err := transformCluster(cluster, i+1)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to transform cluster %s", cluster.Name)
		}
		response[i] = *c
	}

	return response, nil
}

func transformCluster(cluster api.Cluster, index int) (*Cluster, error) {
	response := &Cluster{
		ID:           index,
		Name:         cluster.Name,
		ItemIcons:    clusterItemIcons(cluster),
		Properties:   clusterProperties(cluster),
		Indicators:   clusterIndicators(cluster),
		CanaryChecks: cluster.CanaryChecks,
		Metrics:      cluster.Metrics,
	}

	return response, nil
}

func clusterItemIcons(cluster api.Cluster) []ItemIcon {
	itemIcons := []ItemIcon{}
	return itemIcons
}

func clusterProperties(cluster api.Cluster) []Property {
	properties := []Property{
		{
			ID:      1,
			Service: "cpu",
			Icon:    "cpu",
			Label:   strconv.Itoa(cluster.Kubernetes.CPU),
		},
		{
			ID:      2,
			Service: "memory",
			Icon:    "memory",
			Label:   fmt.Sprintf("%d GB", cluster.Kubernetes.Memory/gb),
		},
		{
			ID:      3,
			Service: "storage",
			Icon:    "disk",
			Label:   fmt.Sprintf("%d GB", cluster.Kubernetes.Disk/gb),
		},
		{
			ID:      5,
			Service: "kernel",
			Icon:    "linux",
			Label:   cluster.Kubernetes.KernelVersion,
			Badge:   len(cluster.Kubernetes.KernelAlerts),
		},
		{
			ID:      6,
			Service: "kubernetes",
			Icon:    "kubernetes",
			Label:   cluster.Kubernetes.KubeletVersion,
			Badge:   len(cluster.Kubernetes.KubeletAlerts),
		},
		{
			ID:      7,
			Service: "os",
			Icon:    "ubuntu",
			Label:   cluster.Kubernetes.OSVersion,
			Badge:   len(cluster.Kubernetes.OSAlerts),
		},
		{
			ID:      8,
			Service: "runtime",
			Icon:    "cargo-ship-50",
			Label:   cluster.Kubernetes.CRIVersion,
			Badge:   len(cluster.Kubernetes.CRIAlerts),
		},
	}

	return properties
}

func clusterIndicators(cluster api.Cluster) []Indicator {
	indicators := []Indicator{
		{
			ID:             "servers",
			IndicatorIcon:  "server",
			IndicatorIcons: serverIndicatorIcons(cluster),
		},
	}

	return indicators
}

func serverIndicatorIcons(cluster api.Cluster) []IndicatorIcon {
	icons := make([]IndicatorIcon, len(cluster.Kubernetes.Nodes))

	for index, node := range cluster.Kubernetes.Nodes {
		icon := IndicatorIcon{
			ID:    node.Name,
			Color: "#336600",
			Shape: "square",
		}

		if len(node.Alerts) > 0 {
			icon.Color = "#990000"
			icon.Overlay = "mdi-bell"
		}

		icons[index] = icon
	}

	return icons
}
