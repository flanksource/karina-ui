package providers

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/karina-ui/pkg/api"
	"github.com/pkg/errors"
	prometheusapi "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type MetricConfig struct {
	Min     string
	Max     string
	Current string
	Total   string
}

var metrics = map[string]MetricConfig{
	"cpu_usage": MetricConfig{
		Min:     "min(sum(rate(node_cpu_seconds_total{mode!=\"idle\",mode!=\"iowait\"}[5m]) * 100) BY (instance))",
		Max:     "max(sum(rate(node_cpu_seconds_total{mode!=\"idle\",mode!=\"iowait\"}[5m]) * 100) BY (instance))",
		Current: "avg(sum(rate(node_cpu_seconds_total{mode!=\"idle\",mode!=\"iowait\"}[5m]) * 100) BY (instance))",
	},
	"cpu_resources": MetricConfig{
		Min:     "sum(kube_pod_container_resource_requests_cpu_cores)",
		Max:     "sum(sum(kube_pod_container_resource_limits_cpu_cores))",
		Current: "sum (rate (container_cpu_usage_seconds_total{id=\"/\"}[5m]))",
		Total:   "sum(kube_node_status_capacity_cpu_cores)",
	},
	"memory": MetricConfig{
		Min:     "min(instance:node_memory_utilisation:ratio * 100)",
		Max:     "max(instance:node_memory_utilisation:ratio * 100)",
		Current: "avg(instance:node_memory_utilisation:ratio * 100)",
	},
	"memory_resources": MetricConfig{
		Min:     "sum(kube_pod_container_resource_requests_memory_bytes)",
		Max:     "sum(kube_pod_container_resource_limits_memory_bytes)",
		Current: "sum(container_memory_working_set_bytes)",
		Total:   "sum(kube_node_status_capacity_memory_bytes)",
	},
	"disk": MetricConfig{
		Min:     "min(100 - ((node_filesystem_avail_bytes{mountpoint=\"/\",fstype!=\"rootfs\"} * 100) / node_filesystem_size_bytes{mountpoint=\"/\",fstype!=\"rootfs\"}))",
		Max:     "max(100 - ((node_filesystem_avail_bytes{mountpoint=\"/\",fstype!=\"rootfs\"} * 100) / node_filesystem_size_bytes{mountpoint=\"/\",fstype!=\"rootfs\"}))",
		Current: "avg(100 - ((node_filesystem_avail_bytes{mountpoint=\"/\",fstype!=\"rootfs\"} * 100) / node_filesystem_size_bytes{mountpoint=\"/\",fstype!=\"rootfs\"}))",
	},
}

var defaultPrometheusDuration = 5 * time.Minute

type Prometheus struct {
	clients map[string]prometheusapi.Client
}

func NewPrometheus(config api.Config) (Provider, error) {
	clients := map[string]prometheusapi.Client{}

	for name, cluster := range config.Clusters {
		transportConfig := prometheusapi.DefaultRoundTripper.(*http.Transport)
		transportConfig.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}

		client, err := prometheusapi.NewClient(prometheusapi.Config{
			Address:      cluster.Prometheus,
			RoundTripper: transportConfig,
		})

		if err != nil {
			return nil, errors.Wrapf(err, "failed to create prometheus client for host %s", cluster.Prometheus)
		}

		clients[name] = client
	}

	prometheus := &Prometheus{clients: clients}
	return prometheus, nil
}

func (p *Prometheus) Fetch(cluster *api.Cluster, config api.ClusterConfiguration) error {
	var err error

	client, _ := p.clients[cluster.Name]
	v1api := v1.NewAPI(client)

	if cluster.Metrics == nil {
		cluster.Metrics = map[string]api.Metric{}
	}

	for name, metricConfig := range metrics {
		metric := api.Metric{}
		metric.MinValue, err = p.getMetric(v1api, metricConfig.Min, defaultPrometheusDuration)
		if err != nil {
			logger.Errorf("failed to get metric for %s/min for cluster %s", name, cluster.Name)
			continue
		}

		metric.MaxValue, err = p.getMetric(v1api, metricConfig.Max, defaultPrometheusDuration)
		if err != nil {
			logger.Errorf("failed to get metric for %s/max for cluster %s", name, cluster.Name)
			continue
		}

		metric.Current, err = p.getMetric(v1api, metricConfig.Current, defaultPrometheusDuration)
		if err != nil {
			logger.Errorf("failed to get metric for %s/current for cluster %s", name, cluster.Name)
			continue
		}

		if metricConfig.Total != "" {
			metric.Total, err = p.getMetric(v1api, metricConfig.Total, defaultPrometheusDuration)
			if err != nil {
				logger.Errorf("failed to get metric for %s/average for cluster %s", name, cluster.Name)
				continue
			}
		}

		cluster.Metrics[name] = metric
	}

	return nil
}

func (p *Prometheus) getMetric(api v1.API, metric string, timeframe time.Duration) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rangeOptions := v1.Range{
		Start: time.Now().Add(-1 * timeframe),
		End:   time.Now(),
		Step:  timeframe,
	}
	result, warnings, err := api.QueryRange(ctx, metric, rangeOptions)

	if err != nil {
		return 0, errors.Wrapf(err, "failed to get metric %s", metric)
	}
	if len(warnings) > 0 {
		logger.Infof("Warnings: %v", warnings)
	}

	// ensure matrix result
	matrix, ok := result.(model.Matrix)
	if !ok {
		logger.Errorf("Result is not a matrix")
		return 0, errors.Errorf("Result is not a matrix")
	}

	if len(matrix) < 1 {
		logger.Errorf("Matrix result is empty")
		return 0, nil
	}

	if len(matrix[0].Values) < 1 {
		logger.Errorf("Matrix values is empty")
		return 0, nil
	}

	return float64(matrix[0].Values[0].Value), nil
}

func (p *Prometheus) Name() string {
	return "prometheus"
}
