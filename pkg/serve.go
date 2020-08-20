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
		resp.Write(json)
	}

}
