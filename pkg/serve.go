package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/flanksource/commons/logger"
	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
  	"gopkg.in/yaml.v2"
  	v1 "k8s.io/api/core/v1"
  	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

		properties = []api.Property{}
				
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

			properties = MergeNode(node, properties)   	
		}

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
	var kernelAlerts, CRIAlerts, K8sAlerts, OSAlerts = []api.Alert{}, []api.Alert{}, []api.Alert{}, []api.Alert{}
	var Alerts []api.Alert

	jsonProp, jsonPropErr := json.Marshal(node.Status)
	if jsonPropErr != nil {
		logger.Errorf("❗ Kube client failed with %s", jsonPropErr)
	}
	if jsonPropErr := json.Unmarshal([]byte(jsonProp), &prop); jsonPropErr != nil {
		logger.Errorf("❗ Failed to unmarshal json %s", jsonPropErr)
	}

	if len(properties) > 0 {

		if properties[4].Name == "Kernel Version" {

			if properties[4].Value != prop.NodeInfo.KernelVersion {
				log := time.Now()
				mesg := "Node: " + prop.NodeInfo.MachineID + "has different kernel version from node: " + properties[3].Value
				kernelAlerts =  append(kernelAlerts, api.Alert {
					Level:	"one",
					Since:	log,
					Message: mesg,
				})
			}
		}

		if properties[5].Name == "CRI Version" {
			
			if properties[5].Value != prop.NodeInfo.ContainerRuntimeVersion {
				log := time.Now()
				mesg := "Node: " + prop.NodeInfo.MachineID + "has different CRI version from node: " + properties[3].Value
				CRIAlerts =  append(CRIAlerts, api.Alert {
					Level:	"two",
					Since:	log,
					Message: mesg,
				})
			}
		}

		if properties[6].Name == "K8S Version" {
			
			if properties[6].Value != prop.NodeInfo.KubeletVersion {
				log := time.Now()
				mesg := "Node: " + prop.NodeInfo.MachineID + "has different K8s version from node: " + properties[3].Value
				K8sAlerts = append(K8sAlerts, api.Alert {
					Level:	"three",
					Since:	log,
					Message: mesg,
				})
			}
		}

		if properties[7].Name == "OS Version" {
			
			if properties[7].Value != prop.NodeInfo.OSImage {
				log := time.Now()
				mesg := "Node: " + prop.NodeInfo.MachineID + "has different OS version from node: " + properties[3].Value
				OSAlerts = append(OSAlerts, api.Alert {
					Level:	"two",
					Since:	log,
					Message: mesg,
				})
			}
		}
	}

	properties = []api.Property {
		{
			Name: "Memory",
			Value: GetMemory(prop.Allocatable.Memory),
			Icon: "memory",
		},
		{
			Name: "CPU",
			Value: prop.Allocatable.Cpu,
			Icon: "cpu",
		},
		{
			Name: "Storage",
			Value: GetStorage(prop.Allocatable.EphStor),
			Icon: "storage",
		},
		{
			Name: "Node",
			Value: prop.NodeInfo.MachineID,
			Icon: "commit",
		},
		{
			Name: "Kernel Version",
			Value: prop.NodeInfo.KernelVersion,
			Icon: "linux",
			Alerts: kernelAlerts,			
		},
		{
			Name: "CRI Version",
			Value: prop.NodeInfo.ContainerRuntimeVersion,
			Icon: GetCRI(prop.NodeInfo.ContainerRuntimeVersion),
			Alerts: CRIAlerts,
		},
		{
			Name: "K8S Version",
			Value: prop.NodeInfo.KubeletVersion,
			Icon: "kubernetes",
			Alerts: K8sAlerts,
		},
		{
			Name: "OS Version",
			Value: prop.NodeInfo.OSImage,
			Icon: GetOSIcon(prop.NodeInfo.OSImage),
			Alerts: OSAlerts,
		},
		{
			Name: "Threats",
			Value: "",
			Icon: "threat",
			Alerts: Alerts,
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
			logger.Errorf("❗ Failed string conversion: %s", err)
		}
		memoryVal = memoryVal * 1024

	} else if strings.Contains(rawMemory, "Mi") == true {

		rawMemory = strings.TrimSuffix(rawMemory, "Mi")

		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed string conversion: %s", err)
		}
		memoryVal = memoryVal * 1024 * 1024

	} else if strings.Contains(rawMemory, "Gi") == true {

		rawMemory = strings.TrimSuffix(rawMemory, "Gi")

		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to string conversion: %s", err)
		}
		memoryVal = memoryVal * 1024 * 1024 * 1024

	} else if strings.Contains(rawMemory, "Ti") == true {

		rawMemory = strings.TrimSuffix(rawMemory, "Ti")

		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to string conversion: %s", err)
		}
		memoryVal = memoryVal * 1024 * 1024 * 1024 * 1024

	} else {
		memoryVal, err = strconv.ParseUint(rawMemory, 10, 64)
		if err != nil {
			logger.Errorf("❗ Failed to string conversion: %s", err)
		}
	}

	memory := humanize.Bytes(memoryVal)
	return memory
}

func GetStorage(rawStorage string) string {

	storageVal, err := strconv.ParseUint(rawStorage, 10, 64)
	if err != nil {
		logger.Errorf("❗ Failed conversion %s", err)
	}
	storage := humanize.Bytes(storageVal)
	return storage
}