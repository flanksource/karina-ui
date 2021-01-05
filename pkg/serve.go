package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/karina-ui/pkg/api"
	"github.com/flanksource/karina-ui/pkg/api/ui"
	"github.com/flanksource/karina-ui/pkg/providers"
	"github.com/flanksource/kommons"
	"github.com/pkg/errors"
	"gopkg.in/flanksource/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var config = api.Config{}

var providersMap = map[string]providers.ProviderFn{
	"canary":     providers.NewCanary,
	"kubernetes": providers.NewKubernetes,
	"prometheus": providers.NewPrometheus,
}

const kubeConfigsLabelSelector = "karina-ui.flanksource.com/kubeconfig=true"
const clusterNameAnnotation = "karina-ui.flanksource.com/cluster-name"

var providersList []providers.Provider

var globalCache []byte = nil
var DevMode = false

func ServeAPI(resp http.ResponseWriter, req *http.Request) {
	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	for name, clusterConfig := range config.Clusters {
		cluster := &api.Cluster{Name: name, Properties: []api.Property{}}

		for _, provider := range providersList {
			if err := provider.Fetch(cluster, clusterConfig); err != nil {
				logger.Errorf("❌ failed to run provider %s for cluster %s: %v", provider.Name(), name, err)
				continue
			}
		}

		clusters = append(clusters, *cluster)
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

func ServeUIAPI(resp http.ResponseWriter, req *http.Request) {
	if DevMode && globalCache != nil {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Write(globalCache)
		return
	}

	logger.Infof("🚀 Fetching data")
	var clusters []api.Cluster
	for name, clusterConfig := range config.Clusters {
		cluster := &api.Cluster{Name: name, Properties: []api.Property{}}

		for _, provider := range providersList {
			if err := provider.Fetch(cluster, clusterConfig); err != nil {
				logger.Errorf("❌ failed to run provider %s for cluster %s: %v", provider.Name(), name, err)
				continue
			}
		}

		clusters = append(clusters, *cluster)
	}

	fc, err := ui.Transform(clusters)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(fc)
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte(err.Error()))
	} else {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Write(json)
		if DevMode {
			globalCache = json
		}
	}
}

func ParseConfiguration(path string) error {
	if path == "" {
		logger.Infof("Config file not provided trying to autodiscover from cluster")
		return AutodiscoverConfiguration()
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return errors.Wrap(err, "failed to unmarshal config file")
	}
	if err := loadProviders(); err != nil {
		return errors.Wrap(err, "failed to load providers")
	}
	return nil
}

func AutodiscoverConfiguration() error {
	ctx := context.Background()

	client, err := AutodiscoverClient()
	if err != nil {
		return errors.Wrap(err, "failed to autodiscover client")
	}

	config.Clusters = map[string]api.ClusterConfiguration{}

	clientset, err := client.GetClientset()
	if err != nil {
		return errors.Wrap(err, "failed to get clientset")
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to list namespaces")
	}

	for _, ns := range namespaces.Items {
		listOptions := metav1.ListOptions{LabelSelector: kubeConfigsLabelSelector}
		secrets, err := clientset.CoreV1().Secrets(ns.Name).List(ctx, listOptions)
		if err != nil {
			logger.Errorf("failed to read secrets in namespace %s: %v", ns.Name, err)
			continue
		}
		for _, secret := range secrets.Items {
			clusterName, found := secret.Annotations[clusterNameAnnotation]
			if !found {
				clusterName = secret.Name
			}
			kubeconfig, found := secret.Data["kubeconfig.yml"]
			if !found {
				logger.Errorf("failed to find kubeconfig.yml in secret %s/%s", ns.Name, secret.Name)
				continue
			}
			cfg, err := DiscoverConfig(kubeconfig)
			if err != nil {
				logger.Errorf("failed to discover config for cluster %s", clusterName)
				continue
			}
			config.Clusters[clusterName] = *cfg
		}
	}

	yml, _ := yaml.Marshal(config)
	fmt.Printf("Autodiscovered config:\n%s\n", yml)

	if err := loadProviders(); err != nil {
		return errors.Wrap(err, "failed to load providers")
	}

	return nil
}

func AutodiscoverClient() (*kommons.Client, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		kubeconfigBytes, err := ioutil.ReadFile(kubeconfig)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read kube config %s", kubeconfig)
		}
		return kommons.NewClientFromBytes(kubeconfigBytes)
	}

	logger.Debugf("Failed to find kube config, trying in cluster config")
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get in cluster config")
	}
	return kommons.NewClient(restConfig, logger.StandardLogger()), nil
}

func DiscoverConfig(kubeconfig []byte) (*api.ClusterConfiguration, error) {
	client, err := kommons.NewClientFromBytes(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create k8s client")
	}

	clientset, err := client.GetClientset()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get clientset")
	}

	canaryChecker, err := getIngress(clientset, "platform-system", "canary-checker")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ingress platform-system/canary-checker")
	}

	prometheus, err := getIngress(clientset, "monitoring", "prometheus")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ingress monitoring/prometheus")
	}

	alertManager, err := getIngress(clientset, "monitoring", "alertmanager")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ingress monitoring/alertmanager")
	}

	tmpf, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create temp file")
	}
	if err := ioutil.WriteFile(tmpf.Name(), kubeconfig, 0600); err != nil {
		return nil, errors.Wrap(err, "failed to write kubeconfig to file")
	}

	config := &api.ClusterConfiguration{
		CanaryChecker: canaryChecker + "/api",
		Prometheus:    prometheus,
		AlertManager:  alertManager,
		Kubeconfig:    tmpf.Name(),
	}
	return config, nil
}

func getIngress(client *kubernetes.Clientset, namespace, name string) (string, error) {
	ctx := context.Background()
	ingress, err := client.ExtensionsV1beta1().Ingresses(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return "", errors.Wrapf(err, "failed to get ingress %s/%s", namespace, name)
	}

	for _, rule := range ingress.Spec.Rules {
		if rule.Host != "" {
			return "https://" + rule.Host, nil
		}
	}

	return "", errors.Errorf("no valid host found for ingress %s/%s", namespace, name)
}

func loadProviders() error {
	providersList = []providers.Provider{}
	for name, providerFn := range providersMap {
		provider, err := providerFn(config)
		if err != nil {
			return errors.Wrapf(err, "failed to create provider %s", name)
		}
		providersList = append(providersList, provider)
	}
	return nil
}
