package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
	"gopkg.in/yaml.v2"
)

var config = make(map[string]api.ClusterConfiguration)

func Serve(resp http.ResponseWriter, req *http.Request) {
	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	for name, cluster := range config {


		var prometheus api.PrometheusData

		prometheusResp, err := net.GET(cluster.Prometheus)
		if err != nil {
			logger.Errorf("❌ Canary Check failed with %s", err)
			continue
		}
	

		if err := json.Unmarshal([]byte(prometheusResp), &prometheus); err != nil {
			logger.Errorf("❌ Failed to unmarshal json %s", err)
			continue
		}

		//fmt.Printf("promdata: %+v\n", prometheus)
		//fmt.Printf("------------------------------------------------------\n\n\n\n\n\n", prometheus)

		//metrics := ["node_memory_MemAvailable_bytes", "machine_cpu_cores", "kube_node_status_allocatable_cpu_cores"]



		

		for _, data := range prometheus.Data {

			metrics := ["node_memory_MemAvailable_bytes", "machine_cpu_cores", "kube_node_status_allocatable_cpu_cores"]

			for metric := range metrics
			

			fmt.Printf("%+v\n\n", data.Metric)


		}


		var canary api.Canarydata
		canaryResp, err := net.GET(cluster.CanaryChecker)
		if err != nil {
			logger.Errorf("❌ Canary Check failed with %s", err)
			continue
		}
		if err := json.Unmarshal([]byte(canaryResp), &canary); err != nil {
			logger.Errorf("❌ Failed to unmarshal json %s", err)
			continue
		}

		clusters = append(clusters, api.Cluster{
			Name: name,
			Properties: []api.Property{
				{
					Name:  "CPU",
					Icon:  "cpu",
					Type:  "cpu",
					Value: "72",
				},
				{
					Name:  "Memory",
					Icon:  "memory",
					Type:  "mem",
					Value: "128",
				},
				{
					Name:  "Disk",
					Icon:  "disk",
					Type:  "disk",
					Value: "100",
				},
			},
			CanaryChecks: canary.Checks,
			Prometheus: prometheus.Data,
		})
	}

	json, err := json.Marshal(clusters)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
	} else {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Write(json)
	}
}

func ParseConfiguration(path string) (map[string]api.ClusterConfiguration, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatalf("❌ Cluster configuration failed with: %v", err)
		return nil, err
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		logger.Fatalf("❌ Cluster configuration failed with: %v", err)
		return nil, err
	}
	return config, err
}
