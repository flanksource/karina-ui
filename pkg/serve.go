package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/flanksource/karina-ui/pkg/api"
)

func Serve(resp http.ResponseWriter, req *http.Request) {

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
					Type:  "memory",
					Value: "128",
				},
				{
					Name:  "Disk",
					Type:  "disk",
					Value: "100",
				},
				{
					Name:  "Commit",
					Type:  "commit",
					Value: "146cfae",
					Alerts: []api.Alert{
						{
							Level:	"string",
							//Since:	time.Time,
							Message:"string",
						},
					},
				},
				{
					Name:  "Kernel",
					Type:  "linux",
					Value: "4.15.0-54-generic",
					Alerts: []api.Alert{
						{
							Level:	"string",
							//Since:	time.Time,
							Message:"string",
						},
						{
							Level:	"string",
							//Since:	time.Time,
							Message:"string",
						},
						{
							Level:	"string",
							//Since:	time.Time,
							Message:"string",
						},
					},
				},
				{
					Name:  "Kubernetes",
					Type:  "kubernetes",
					Value: "v1.16.3",
					Alerts: []api.Alert{
						{
							Level:	"string",
							//Since:	time.Time,
							Message:"string",
						},
						{
							Level:	"string",
							//Since:	2012-04-23T18:25:43.511Z,
							Message:"string",
						},
					},
				},
				{
					Name:  "OSVersion",
					Type:  "ubuntu",
					Value: "18.04 LTS",
					Alerts: []api.Alert{
						{
							Level:	"string",
							//Since:	time.Time,
							Message:"string",
						},
						{
							Level:	"string",
							//Since:	2012-04-23T18:25:43.511Z,
							Message:"string",
						},
					},
				},
			},
			CanaryChecks: []api.CanaryCheck{
				{
					
					Name:		"string",
					Namespace: 	"string",
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",
					Uptime1H:	"string",
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,
							//Timestamp:	null,
							Message:	"string",
						},
					},
				},
			},
			Nodes: []api.Node{
				{
					Name:   "string",
					IP:		"string",
					Alerts: []api.Alert{
						{
							Level:	"string",
							//Since:	time.Time,
							Message:"string",
						},
					},
				},
			},
			Alerts: []api.Alert{
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
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