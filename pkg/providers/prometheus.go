package providers

import (
	"context"
	"crypto/tls"
	"net/http"
	"strconv"
	"time"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/karina-ui/pkg/api"
	"github.com/pkg/errors"
	prometheusapi "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

var metrics = map[string]string{
	"cpu":    "sum(rate(node_cpu_seconds_total{mode!=\"idle\",mode!=\"iowait\"}[5m]) * 100) BY (instance)",
	"memory": "instance:node_memory_utilisation:ratio * 100",
	"disk":   "100 - ((node_filesystem_avail_bytes{mountpoint=\"/\",fstype!=\"rootfs\"} * 100) / node_filesystem_size_bytes{mountpoint=\"/\",fstype!=\"rootfs\"})",
}

var defaultPrometheusDuration = 5 * time.Minute

type Timeseries struct {
	Time       float64 `json:"time"`
	FloatValue float64 `json:"floatValue"`
	Value      string  `json:"value"`
}

type InstanceTimeseries map[string]Timeseries

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
	client, _ := p.clients[cluster.Name]
	v1api := v1.NewAPI(client)

	if cluster.Metrics == nil {
		cluster.Metrics = map[string]api.Metric{}
	}

	for name, metric := range metrics {
		instanceTimeseries, err := p.getMetric(v1api, metric, defaultPrometheusDuration)
		if err != nil {
			logger.Errorf("failed to get metric for %s for cluster %s", name, cluster.Name)
			continue
		}

		var count = len(instanceTimeseries)
		var sum float64 = 0

		for _, ts := range instanceTimeseries {
			sum += ts.FloatValue
		}

		medium := sum / float64(count)

		cluster.Metrics[name] = api.Metric{
			FloatValue: medium,
			Value:      strconv.FormatFloat(medium, 'f', 1, 64) + "%",
		}
	}

	return nil
}

func (p *Prometheus) getMetric(api v1.API, metric string, timeframe time.Duration) (InstanceTimeseries, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rangeOptions := v1.Range{
		Start: time.Now().Add(-1 * timeframe),
		End:   time.Now(),
		Step:  5 * time.Minute,
	}
	result, warnings, err := api.QueryRange(ctx, metric, rangeOptions)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get metric %s", metric)
	}
	if len(warnings) > 0 {
		logger.Infof("Warnings: %v", warnings)
	}

	// ensure matrix result
	matrix, ok := result.(model.Matrix)
	if !ok {
		logger.Errorf("Result is not a matrix")
		return nil, errors.Errorf("Result is not a matrix")
	}

	if len(matrix) < 1 {
		logger.Errorf("Matrix result is empty")
		return InstanceTimeseries{}, nil
	}

	results := InstanceTimeseries{}

	for _, instanceMetric := range matrix {
		instanceName := string(instanceMetric.Metric["instance"])
		result := Timeseries{}
		// TODO (toni) Take aggregate here ?
		for _, t := range instanceMetric.Values {
			result = Timeseries{
				Time:       float64(t.Timestamp.Unix()),
				FloatValue: float64(t.Value),
				Value:      t.Value.String(),
			}
		}
		results[instanceName] = result
	}

	return results, nil
}

func (p *Prometheus) Name() string {
	return "prometheus"
}
