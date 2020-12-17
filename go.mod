module github.com/flanksource/karina-ui

go 1.14

require (
	github.com/flanksource/canary-checker v0.11.7
	github.com/flanksource/commons v1.4.0
	github.com/go-yaml/yaml v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	gopkg.in/flanksource/yaml.v3 v3.1.1
)

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

replace k8s.io/client-go => k8s.io/client-go v0.17.7
