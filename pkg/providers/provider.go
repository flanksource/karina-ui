package providers

import "github.com/flanksource/karina-ui/pkg/api"

type Provider interface {
	Fetch(cluster *api.Cluster, config api.ClusterConfiguration) error
	Name() string
}

type ProviderFn func(api.Config) (Provider, error)
