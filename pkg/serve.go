package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
	"github.com/pkg/errors"
	"gopkg.in/flanksource/yaml.v3"
)

var config = make(map[string]api.ClusterConfiguration)

func Serve(resp http.ResponseWriter, req *http.Request) {
	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	for name, cluster := range config {

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
	if path == "" {
		return nil, errors.Errorf("config file flag is missing")
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config file")
	}
	return config, nil
}
