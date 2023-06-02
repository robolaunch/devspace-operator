# Installation

## Prerequisites

For robolaunch DevSpace Operator, these prerequisites should be satisfied:

|     Tool     |       Version      |
|:------------:|:------------------:|
|  Kubernetes  |  `v1.21` and above |
| Cert-Manager | `v1.8.x` and above |
|    OpenEBS   | `v3.x.x` and above |

### Labeling Node

Select an active node from your cluster and add these labels:

```bash
kubectl label <NODE> robolaunch.io/organization=robolaunch
kubectl label <NODE> robolaunch.io/team=development
kubectl label <NODE> robolaunch.io/region=europe-east
kubectl label <NODE> robolaunch.io/cloud-instance=cluster
kubectl label <NODE> robolaunch.io/cloud-instance-alias=cluster-alias
```

## Installing DevSpace Operator

### via Helm

Add robolaunch Helm repository and update:

```bash
helm repo add robolaunch https://robolaunch.github.io/charts/
helm repo update
```

Install latest version of DevSpace Operator (remove `--devel` for getting latest stable version):

```bash
helm upgrade -i devspace-operator robolaunch/devspace-operator  \
--namespace devspace-system \
--create-namespace \
--devel
```

Or you can specify a version (remove the `v` letter at the beginning of the release or tag name):

```bash
VERSION="0.1.0-alpha.1"
helm upgrade -i devspace-operator robolaunch/devspace-operator  \
--namespace devspace-system \
--create-namespace \
--version $VERSION
```

### via Manifest

Deploy DevSpace Operator one-file YAML using the command below:

```bash
# select a tag
TAG="v0.1.0-alpha.1"
kubectl apply -f https://raw.githubusercontent.com/robolaunch/devspace-operator/$TAG/hack/deploy/manifests/devspace_operator.yaml
```

## Uninstalling DevSpace Operator

To uninstall DevSpace Operator installed with Helm, run the following commands:

```bash
helm delete devspace-operator -n devspace-system
kubectl delete ns devspace-system
```

To uninstall DevSpace Operator installed with one-file YAML, run the following commands:
```bash
# find the tag you installed
TAG="v0.1.0-alpha.1"
kubectl delete -f https://raw.githubusercontent.com/robolaunch/devspace-operator/$TAG/hack/deploy/manifests/devspace_operator.yaml
```