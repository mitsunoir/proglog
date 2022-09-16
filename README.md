# Proglog

## Deploy to local Kubernetes using Kind

1. Install kubectl command

   Command line tool for Kubernetes.

2. Install Kind

   Kubernetes cluster execution tool using Docker containers as nodes, which suits to local development, test and CI.

3. Create cluster

```
kind create cluster
```

You can see the result with:

```
kubectl cluster-info
```

4. Load Docker image

```
kind load docker-image github.com/mitsunoir/proglog:0.0.1
```

5. Install Helm

   Kubernetes package manager.

6. Install Helm chart

```
helm install proglog deploy/proglog
```
