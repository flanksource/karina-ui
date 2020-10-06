package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"reflect"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
  	"gopkg.in/yaml.v2"
  	v1 "k8s.io/api/core/v1"
  	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"	
)

var config = make(map[string]api.ClusterConfiguration)
var clientset *kubernetes.Clientset
var properties *api.Properties
var freshproperties api.Properties

var name string

func Serve(resp http.ResponseWriter, req *http.Request) {

	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	var canary api.Canarydata

	for name, cluster := range config {

		properties = &api.Properties{}
		
		clientset, clientsetErr := GetClient(cluster.Kubeconfig)
		if clientsetErr != nil {
			logger.Errorf("❗ Get K8s client failed with %s", clientsetErr)
			continue
		}

		nodeList, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
		   logger.Errorf("❗ Get K8s client properties failed with %s", err)
			continue
		}
		

		for _, node := range nodeList.Items {

		    properties := MergeNode(node, *properties, name)
		    fmt.Printf("properties: %v", properties)
			
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
				Properties: properties,
	      		CanaryChecks: canary.Checks,
				Nodes: []api.Node {
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

func MergeNode(node v1.Node, properties api.Properties, name string) api.Properties {

	jsonProp, jsonPropErr := json.Marshal(node.Status)
	if jsonPropErr != nil {
		logger.Errorf("❗ Kube client failed with %s", jsonPropErr)
	}
	if jsonPropErr := json.Unmarshal([]byte(jsonProp), &properties); jsonPropErr != nil {
		logger.Errorf("❗ Failed to unmarshal json %s", jsonPropErr)
	}

	if properties.NodeInfo.KernelVersion != node.Status.NodeInfo.KernelVersion {
		properties.Alerts = append(freshproperties.Alerts, api.Alert{
			Level:	"string",
			//Since:	time.Time,
			Message:"string",
			Type: "Kernel Alert",
			Cluster: name,
		})
	}

	//fmt.Printf("in merge node kernelalerts:\n %v\n", freshproperties.Alerts)
	//fmt.Printf("----------------------------------------------\n")

	/*
	if freshproperties.NodeInfo.OsImage != node.Status.NodeInfo.OSImage {
		fmt.Println("Raise OS image alert")
	}
	if freshproperties.NodeInfo.ContainerRuntimeVersion != node.Status.NodeInfo.ContainerRuntimeVersion {
		fmt.Println("Raise Container alert")
	}
	if freshproperties.NodeInfo.KubeletVersion != node.Status.NodeInfo.KubeletVersion {
		fmt.Println("Raise Kubelet alert")
	}
	if freshproperties.NodeInfo.KubeProxyVersion != node.Status.NodeInfo.KubeProxyVersion {
		fmt.Println("Raise Kube Proxy alert")
	}
	if freshproperties.NodeInfo.OperatingSystem != node.Status.NodeInfo.OperatingSystem {
		fmt.Println("Raise OS alert")
	}
	if freshproperties.NodeInfo.Architecture != node.Status.NodeInfo.Architecture {
		fmt.Println("Raise OS arch alert")
	}
	*/

	return properties  
}