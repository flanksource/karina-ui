module github.com/flanksource/karina-ui

go 1.14

require (
	github.com/flanksource/canary-checker v0.11.7
	github.com/flanksource/commons v1.4.0
	//

	github.com/go-yaml/yaml v2.1.0+incompatible // indirect

	// issue-36
	//	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/apimachinery v0.17.7
	k8s.io/client-go v11.0.0+incompatible
	//	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	//	k8s.io/api v0.17.7
	//	k8s.io/apimachinery v0.17.7
	//	k8s.io/client-go v11.0.0+incompatible
	sigs.k8s.io/yaml v1.1.0

)

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

replace k8s.io/client-go => k8s.io/client-go v0.17.7
