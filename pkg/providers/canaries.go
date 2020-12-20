package providers

import (
	"encoding/json"

	"github.com/flanksource/commons/net"
	"github.com/flanksource/karina-ui/pkg/api"
	"github.com/pkg/errors"
)

type Canary struct {
}

func NewCanary(config api.Config) (Provider, error) {
	return &Canary{}, nil
}

func (c *Canary) Fetch(cluster *api.Cluster, config api.ClusterConfiguration) error {
	var canary api.Canarydata
	canaryResp, err := net.GET(config.CanaryChecker)
	if err != nil {
		return errors.Wrapf(err, "failed to get canary checker response for cluster %s", cluster.Name)
	}
	if err := json.Unmarshal([]byte(canaryResp), &canary); err != nil {
		return errors.Wrapf(err, "failed to unmarshal canary checker response for cluster %s", cluster.Name)
	}

	cluster.CanaryChecks = canary.Checks

	return nil
}

func (c *Canary) Name() string {
	return "canary"
}
