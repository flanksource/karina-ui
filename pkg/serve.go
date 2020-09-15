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
			CanaryChecks: []api.CanaryCheck{
				{
					
					Name:		"dns", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},
				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},
				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},
				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"lan", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		false,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},


				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},
				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"lan", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		false,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},


				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
				},

				{
					
					Name:		"ldap", 				//NAME
					Namespace: 	"platform-system", 	//NAMESPACE
					Type:		"string",
					Endpoint:  	"string",
					Latency1H:	"string",			//LATENCY 1H
					Uptime1H:	"0/2 (0%)",			//UPTIME 1H
					Status: 	[]api.CanaryCheckStatus{
						{
							Pass: 		true,		//STATUS
							//Timestamp:	null, 	//LAST CHECK INTERVAL LAST TRANSITIONED
							Message:	"string",	//MESSAGE
						},
					},
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
