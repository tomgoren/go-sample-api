# Overview

## Build and test
```
cd src/
go build
go test
go run main.go
```

In a new terminal:
```
./test_script.sh
```

## Publish Image
```
podman build . -t ghcr.io/tomgoren/go-sample-api:latest
echo "${GITHUB_TOKEN}" | podman login ghcr.io -u tomgoren --password-stdin
podman push ghcr.io/tomgoren/go-sample-api:latest
```

## Deploy locally

### Bring up podman and kind cluster
```
podman machine stop
podman machine set --rootful
podman machine start

export DOCKER_HOST='unix:///Users/tomgoren/.local/share/containers/podman/machine/qemu/podman.sock'

kind create cluster --name tomtest-1

# kubectx should be set
```

### Apply manifests

```
kubectl apply -f deploy/namespace.yaml
kubectl apply -f deploy/deployment.yaml
kubectl apply -f deploy/service.yaml
```

### Port forward
```
kubectl -n go-api port-forward service/go-api-service 8080:80
```

In a new terminal:
```
./test_script.sh
```
