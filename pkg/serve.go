package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/karina-ui/pkg/api"
	"github.com/flanksource/karina-ui/pkg/api/ui"
	"github.com/flanksource/karina-ui/pkg/providers"
	"github.com/pkg/errors"
	"gopkg.in/flanksource/yaml.v3"
)

var config = api.Config{}

var providersMap = map[string]providers.ProviderFn{
	"canary":     providers.NewCanary,
	"kubernetes": providers.NewKubernetes,
	"prometheus": providers.NewPrometheus,
}

var providersList []providers.Provider

var globalCache []byte = nil

func ServeAPI(resp http.ResponseWriter, req *http.Request) {
	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	for name, clusterConfig := range config.Clusters {
		cluster := &api.Cluster{Name: name, Properties: []api.Property{}}

		for _, provider := range providersList {
			if err := provider.Fetch(cluster, clusterConfig); err != nil {
				logger.Errorf("❌ failed to run provider %s for cluster %s", provider.Name(), name)
				continue
			}
		}

		clusters = append(clusters, *cluster)
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

func ServeUIAPI(resp http.ResponseWriter, req *http.Request) {
	if globalCache != nil {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Write(globalCache)
		return
	}

	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	for name, clusterConfig := range config.Clusters {
		cluster := &api.Cluster{Name: name, Properties: []api.Property{}}

		for _, provider := range providersList {
			if err := provider.Fetch(cluster, clusterConfig); err != nil {
				logger.Errorf("❌ failed to run provider %s for cluster %s", provider.Name(), name)
				continue
			}
		}

		clusters = append(clusters, *cluster)
	}

	fc, err := ui.Transform(clusters)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(fc)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
	} else {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Write(json)
		globalCache = json
	}
}

func ParseConfiguration(path string) error {
	if path == "" {
		return errors.Errorf("config file flag is missing")
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return errors.Wrap(err, "failed to unmarshal config file")
	}
	if err := loadProviders(); err != nil {
		return errors.Wrap(err, "failed to load providers")
	}
	return nil
}

func loadProviders() error {
	providersList = []providers.Provider{}
	for name, providerFn := range providersMap {
		provider, err := providerFn(config)
		if err != nil {
			return errors.Wrapf(err, "failed to create provider %s", name)
		}
		providersList = append(providersList, provider)
	}
	return nil
}
