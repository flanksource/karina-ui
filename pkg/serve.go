package pkg

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"github.com/flanksource/karina-ui/pkg/api"





	//"context"
	"fmt"
	//"io/ioutil"
	//"net/http"
	//"sort"
	//"strings"
	//"time"

	//"github.com/flanksource/canary-checker/api/external"
	//"k8s.io/api/extensions/v1beta1"

	//"k8s.io/apimachinery/pkg/util/intstr"

	//canaryv1 "github.com/flanksource/canary-checker/api/v1"
	//"github.com/flanksource/canary-checker/pkg"
	//"github.com/flanksource/canary-checker/pkg/metrics"
	"github.com/flanksource/commons/logger"
	//perrors "github.com/pkg/errors"
	//v1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"sigs.k8s.io/yaml"

	//"golang.org/x/sync/semaphore"
	"k8s.io/client-go/kubernetes"
)


const (
	podLabelSelector   = "canary-checker.flanksource.com/podName"
	podCheckSelector   = "canary-checker.flanksource.com/podCheck"
	podGeneralSelector = "canary-checker.flanksource.com/generated"
)

type PodChecker struct {
	//lock *semaphore.Weighted
	k8s  *kubernetes.Clientset
	//ng   *NameGenerator

	latestNodeIndex int
}

var c *PodChecker



func Serve(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("1\n")

	fmt.Println("2\n")

	logger.Debugf("Running pod check %s")
	fmt.Println("3\n")
	five := int64(5)
	fmt.Println("4\n")

	a := &five
	fmt.Printf("a-&five: %v\n",a)
	b := metav1.ListOptions{TimeoutSeconds: &five}
	fmt.Printf("b-meta...: %v\n", b)


	fmt.Printf("c wacho: %v\n", c)
	if c != nil {
		nodes, err := c.k8s.CoreV1().Nodes().List(metav1.ListOptions{TimeoutSeconds: &five})
		fmt.Println(err)
		fmt.Println(nodes)
	} else{
		fmt.Println(c)

	}
	
	fmt.Println("5\n")

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