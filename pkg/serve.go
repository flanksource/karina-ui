package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
	"gopkg.in/yaml.v2"
)


var config = make(map[string]ClusterConfiguration)
var clusters = []api.Cluster{}

type ClusterConfiguration struct {
	CanaryChecker string `yaml:"canaryChecker,omitempty"`
	Prometheus    string `yaml:"prometheus,omitempty"`
	AlertManager  string `yaml:"alertmanager,omitempty"`
	// A kubeconfig with connection details to a kubernetes cluster
	Kubeconfig string `yaml:"kubeconfig,omitempty"`
}

func ParseConfiguration(path string) (map[string]ClusterConfiguration, error) {
	
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

//On startup you pass the config file path as a cobra argument,
//parse it and
//then store it in global variable in the Serve method, 
//you iterate over the map


func Serve(resp http.ResponseWriter, req *http.Request) {

	var clusterConfig, erra = ParseConfiguration("pkg/config/kubeconfig.yml")

	//To handle
	if erra == nil {
		fmt.Printf("\n")
	}

    keys := reflect.ValueOf(clusterConfig).MapKeys()

    strkeys := make([]string, len(keys))
    for i := 0; i < len(keys); i++ {
        strkeys[i] = keys[i].String()
    }

    for x := 0; x < len(strkeys); x++{
    	var clusterKey = strkeys[x]
    	var clusterValues = clusterConfig[clusterKey]

    	var canaryUrl = clusterValues.CanaryChecker
    	//var prometheusUrl = clusterValues.Prometheus
    	//var alertManager = clusterValues.AlertManager
    	//var kubeConfig = clusterValues.Kubeconfig
    	
/*    	fmt.Printf("%d\n", x)
    	fmt.Printf("%v\n", canaryUrl)
    	fmt.Printf("\n")
*/
		var canary api.Canarydata
		var respCanary, err = net.GET(canaryUrl)
		json.Unmarshal([]byte(respCanary), &canary)

	    if err != nil {
	      	logger.Infof("Canary Check failed with %s\n", err)
	    }

	    name := "Cluster" + strconv.Itoa(x+1)

	    logger.Infof("Retrieving data...")
	    
		clusters = append(clusters, api.Cluster{
			Name: name,
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