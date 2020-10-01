package pkg

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
  	"gopkg.in/yaml.v2"
  	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"	
)

var clientset *kubernetes.Clientset
var config = make(map[string]api.ClusterConfiguration)

func Serve(resp http.ResponseWriter, req *http.Request) {
	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster

	for name, cluster := range config {

		clientset, clientsetErr := GetClient(cluster.Kubeconfig)
		if clientsetErr != nil {
			logger.Errorf("❗ Get K8s client failed with %s", clientsetErr)
			continue
		}

		clientProp, clientPropErr := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		if clientPropErr != nil {
			logger.Errorf("❗ Get K8s client properties failed with %s", clientPropErr)
			continue
		}

		var clusterProp api.Properties
		//convert type *"k8s.io/api/core/v1".NodeList --> type api.Properties
		jsonProp, jsonPropErr := json.Marshal(clientProp.Items[0].Status)

		if jsonPropErr != nil {
			logger.Errorf("❗ Kube client failed with %s", jsonPropErr)
			continue
		}
		if jsonPropErr := json.Unmarshal([]byte(jsonProp), &clusterProp); jsonPropErr != nil {
			logger.Errorf("❗ Failed to unmarshal json %s", jsonPropErr)
			continue
		}	

		var canary api.Canarydata
		canaryResp, err := net.GET(cluster.CanaryChecker)
		if err != nil {
			logger.Errorf("❗ Canary Check failed with %s", err)
			continue
		}
		if err := json.Unmarshal([]byte(canaryResp), &canary); err != nil {
			logger.Errorf("❗ Failed to unmarshal json %s", err)
			continue
		}

		clusters = append(clusters, api.Cluster{
			Name: name,
			Properties: clusterProp,
      		CanaryChecks: canary.Checks,
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

func GetClient(configPath string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}