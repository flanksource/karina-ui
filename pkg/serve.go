package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"reflect"
	"strings"

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
var properties []api.Property
var freshproperties []api.Property


func Serve(resp http.ResponseWriter, req *http.Request) {

	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	var canary api.Canarydata

	for name, cluster := range config {

		//properties = []api.Property{}

		fmt.Printf("start: %v\n\n", name)
		
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

		    properties := MergeNode(node, properties)
		    	
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

			fmt.Printf("end: %v\n\n", name)
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

func MergeNode(node v1.Node, properties []api.Property) []api.Property {

	var prop api.Properties




	jsonProp, jsonPropErr := json.Marshal(node.Status)
	if jsonPropErr != nil {
		logger.Errorf("❗ Kube client failed with %s", jsonPropErr)
	}
	if jsonPropErr := json.Unmarshal([]byte(jsonProp), &prop); jsonPropErr != nil {
		logger.Errorf("❗ Failed to unmarshal json %s", jsonPropErr)
	}


	fmt.Printf("old properties in merge method: %v\n\n", properties)
	fmt.Printf("new properties: %v\n\n", prop)


	/*if properties.NodeInfo.KernelVersion == node.Status.NodeInfo.KernelVersion {
		fmt.Printf("Raise Kernel version alert %v\n", name)
		properties = append(properties, api.Alert{
			Level:	"string",
			//Since:	time.Time,
			Message:"string",
			Type: "Kernel Alert",
			Cluster: name,
		})
	}

*/


	properties = []api.Property {
		{
			Name: "Memory",
			Value: prop.Allocatable.Memory,
			Icon: "memory",
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "CPU",
			Value: prop.Allocatable.Cpu,
			Icon: "cpu",
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "Storage",
			Value: prop.Allocatable.EphStor,
			Icon: "storage",
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "Node",
			Value: prop.NodeInfo.MachineID,
			Icon: "commit",
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "Kernel",
			Value: prop.NodeInfo.KernelVersion,
			Icon: "linux",
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "K8S Version",
			Value: prop.NodeInfo.KubeletVersion,
			Icon: "kubernetes",
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "OS Version",
			Value: prop.NodeInfo.OSImage,
			Icon: GetOSIcon(prop.NodeInfo.OSImage),
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "Threats",
			Value: "",
			Icon: "threat",
			Alerts:  []api.Alert {
				{
					Level:	"string",
					//Since:	time.Time,
					Message:"string",
				},
			},
		},
		{
			Name: "Updates",
			Value: "", 
			Icon: "safe",
		},
	}

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

	//fmt.Printf("----------------------------------------------\n")

	

	return properties  
}

func GetOSIcon(image string) string {
	icon := "null"
	if strings.Contains(strings.ToLower(image), "ubuntu") == true {
		icon = "ubuntu"
	}
	return icon
}