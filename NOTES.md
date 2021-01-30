# Local Testing Steps

## Karina

spin up cluster and deploy

```bash
karina karina ca generate --name root-ca --cert-path .certs/root-ca.crt --private-key-path .certs/root-ca.key --password foobar  --expiry 1
karina ca generate --name ingress-ca --cert-path .certs/ingress-ca.crt --private-key-path .certs/ingress-ca.key --password foobar  --expiry 1
karina provision kind-cluster -c cluster.yaml
karina deploy phases --base --dex --calico -c cluster.yaml
karina deploy all -c cluster.yaml
```
Configure in-cluster secret:

```bash
kind get kubeconfig --name test-cluster > test-kubeconfig.yaml
kubectl create secret generic test-secret --from-file ./kubeconfig.yml
kubectl annotate secret test-secret karina-ui.flanksource.com/cluster-name=test-k8s
kubectl label secret test-secret karina-ui.flanksource.com/kubeconfig=true
```

## Karina-ui

build Vue assets:

```bash
npm install
npm run build
```

run app:

```bash
KUBECONFIG=`pwd`/kubeconfig.yml go run main.go serve --dev
```
