package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"reflect"
	"strings"
	"strconv"

	"github.com/dustin/go-humanize"
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

	properties = []api.Property {
		{
			Name: "Memory",
			Value: GetMemory(prop.Allocatable.Memory),
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
			Value: GetStorage(prop.Allocatable.EphStor),
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
			Name: "CRI Version",
			Value: prop.NodeInfo.ContainerRuntimeVersion,
			Icon: GetCRI(prop.NodeInfo.ContainerRuntimeVersion),
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
	return properties  
}



func GetCRI(cri string) string {

	icon := ""
	if strings.Contains(strings.ToLower(cri), "containerd") == true {
		icon = "containerd"
	}
	return icon
}

func GetOSIcon(image string) string {

	icon := ""
	if strings.Contains(strings.ToLower(image), "ubuntu") == true {
		icon = "ubuntu"
	}
	return icon
}

func GetMemory(rawMemory string) string {

	var memoryVal uint64
	var err error

	if strings.Contains(rawMemory, "Ki") == true {

		rawMemory = strings.TrimSuffix(rawMemory, "Ki")

		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to srting conversion: %s", err)
		}
		memoryVal = memoryVal * 1024

	} else if strings.Contains(rawMemory, "Mi") == true {

		rawMemory = strings.TrimSuffix(rawMemory, "Mi")

		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to srting conversion: %s", err)
		}
		memoryVal = memoryVal * 1024 * 1024

	} else if strings.Contains(rawMemory, "Gi") == true {

		rawMemory = strings.TrimSuffix(rawMemory, "Gi")

		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to srting conversion: %s", err)
		}
		memoryVal = memoryVal * 1024 * 1024 * 1024

	} else if strings.Contains(rawMemory, "Ti") == true {

		rawMemory = strings.TrimSuffix(rawMemory, "Ti")

		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to srting conversion: %s", err)
		}
		memoryVal = memoryVal * 1024 * 1024 * 1024 * 1024

	} else {
		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to srting conversion: %s", err)
		}
	}

	memory := humanize.Bytes(memoryVal)
	return memory
}


func GetStorage(rawStorage string) string {

	storageVal, err := strconv.ParseUint(rawStorage, 10, 64)
	if err != nil {
		logger.Errorf("❗ Failed to convert %s", err)
	}
	storage := humanize.Bytes(storageVal)
	return storage
}