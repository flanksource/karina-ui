package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
)


func Serve(resp http.ResponseWriter, req *http.Request) {

	var canary api.Canarydata
	var respCanary, err = net.GET("https://canaries.hetzner.lab.flanksource.com/api")
	json.Unmarshal([]byte(respCanary), &canary)
	
    if err != nil {
      	logger.Infof("Canary Check failed with %s\n", err)
    }
    
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


