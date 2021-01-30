# Karina - The Kubernetes Dashboard

This is a GUI for Karina

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

## Local Development Steps

One local development setup is to setup a local cluster to connect to, to run the golang API server and 
to separately run the npm in local live reload:

### Karina

Spin up cluster and deploy:

```bash
karina ca generate --name root-ca --cert-path .certs/root-ca.crt --private-key-path .certs/root-ca.key --password foobar  --expiry 1
karina ca generate --name ingress-ca --cert-path .certs/ingress-ca.crt --private-key-path .certs/ingress-ca.key --password foobar  --expiry 1
karina provision kind-cluster -c cluster.yaml
karina deploy phases --base --dex --calico -c cluster.yaml
karina deploy all -c cluster.yaml
```
Configure in-cluster secret:

```bash
kind get kubeconfig --name test-cluster > kubeconfig.yml
export KUBECONFIG=`pwd`/kubeconfig.yml
kubectl create secret generic test-secret --from-file ./kubeconfig.yml
kubectl annotate secret test-secret karina-ui.flanksource.com/cluster-name=test-k8s
kubectl label secret test-secret karina-ui.flanksource.com/kubeconfig=true
```

### Karina-ui GUI

build Vue assets:

```bash
npm install
npm run build
```

Run local development:

```bash
npm run serve
```

### Karina-ui Backend

run app:

```bash
KUBECONFIG=`pwd`/kubeconfig.yml go run main.go serve --dev
```

