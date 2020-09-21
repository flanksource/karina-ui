package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"gopkg.in/yaml.v2"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
)


type ClusterConfiguration struct {
	CanaryChecker string `yaml:"canaryChecker,omitempty"`
	Prometheus    string `yaml:"prometheus,omitempty"`
	AlertManager  string `yaml:"alertmanager,omitempty"`
	// A kubeconfig with connection details to a kubernetes cluster
	Kubeconfig string `yaml:"kubeconfig,omitempty"`
}

func ParseConfiguration(path string) (map[string]ClusterConfiguration, error) {
	config := make(map[string]ClusterConfiguration)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return nil, err
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("yaml error: %s\n", err)
		return nil, err
	}

	return config, err
}


//You would use that snipped to create map of clusters based on a yaml file:

//On startup you pass the config file path as a cobra argument, parse it and
//then store it in global variable in the Serve method, you iterate over the map


func Serve(resp http.ResponseWriter, req *http.Request) {

	var x, y = ParseConfiguration("pkg/config/kubeconfig.yml")

	if y == nil {
		fmt.Printf("y:\n %s\n", y)
	}
	fmt.Printf("x:\n %s\n", x)

	//for 
	

	var canary api.Canarydata

	var respCanary, err = net.GET("https://canaries.hetzner.lab.flanksource.com/api")
	json.Unmarshal([]byte(respCanary), &canary)
	
    if err != nil {
      	logger.Infof("Canary Check failed with %s\n", err)
    }

    logger.Infof("Retrieving data...")

    
	var clusters = []api.Cluster{
		{
			Name: "cluster01",
			Properties: []api.Property{
				{
					Name:  "CPU",
					Type:  "cpu",
					Value: "72",
				},
				{
					Name:  "Memory",
					Type:  "mem",
					Value: "128",
				},
				{
					Name:  "Disk",
					Type:  "disk",
					Value: "100",
				},
			},
			CanaryChecks: canary.Checks,
		},
		{
			Name: "cluster02",
			Properties: []api.Property{
				{
					Name:  "CPU",
					Type:  "cpu",
					Value: "72",
				},
				{
					Name:  "Memory",
					Type:  "mem",
					Value: "128",
				},
				{
					Name:  "Disk",
					Type:  "disk",
					Value: "100",
				},
			},
			CanaryChecks: canary.Checks,
		},
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


