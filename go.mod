module github.com/flanksource/karina-ui

go 1.14

require (
	github.com/Masterminds/semver v1.5.0
	github.com/flanksource/canary-checker v0.11.7
	github.com/flanksource/commons v1.4.3
	github.com/flanksource/kommons v0.1.6
	github.com/go-yaml/yaml v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	golang.org/x/mod v0.1.1-0.20191105210325-c90efee705ee
	gopkg.in/flanksource/yaml.v3 v3.1.1
	k8s.io/api v0.19.3
	k8s.io/apimachinery v0.19.3
)

replace (
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20191230161307-f3c370f40bfb
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/client-go => k8s.io/client-go v0.19.3
)
