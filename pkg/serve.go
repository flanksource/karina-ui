package pkg

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
	"gopkg.in/yaml.v2"
)


var config = make(map[string]api.ClusterConfiguration)


func Serve(resp http.ResponseWriter, req *http.Request) {
	logger.Infof("🚀 Fetching data")
	var name string
	var clusters = []api.Cluster{}
	var clusterConfig, clusterConfigErr = ParseConfiguration("pkg/config/kubeconfig.yml")

	if clusterConfigErr != nil {
		logger.Fatalf("❌ Cluster configuration failed with: %v", clusterConfigErr)
	}

    keys := reflect.ValueOf(clusterConfig).MapKeys()
    if len(keys) < 1 {
    	logger.Infof("⚠️  Cluster configuration has no data")
    }

    strKeys := make([]string, len(keys))
    for i := 0; i < len(keys); i++ {
        strKeys[i] = keys[i].String()
    }

    for x := 0; x < len(strKeys); x++ {
    	var clusterKey = strKeys[x]
    	var clusterValues = clusterConfig[clusterKey]

    	var canaryUrl = clusterValues.CanaryChecker    	
		var canary api.Canarydata
		var canaryResp, err = net.GET(canaryUrl)
		json.Unmarshal([]byte(canaryResp), &canary)
	    if err != nil {
	      	logger.Fatalf("❌ Canary Check failed with %s\n", err)
	    }


	    //var prometheusUrl = clusterValues.Prometheus
    	//var alertManager = clusterValues.AlertManager
    	//var kubeConfig = clusterValues.Kubeconfig

	    if x<10 {
	    	name = "Cluster0" + strconv.Itoa(x+1)
	    } else {
	    	name = "Cluster" + strconv.Itoa(x+1)
	    }
	    
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


func ParseConfiguration(path string) (map[string]api.ClusterConfiguration, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return config, err
}
