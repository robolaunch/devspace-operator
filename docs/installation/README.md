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
kubectl label <NODE> robolaunch.io/team=robotics
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
helm upgrade -i robot-operator robolaunch/robot-operator  \
--namespace robot-system \
--create-namespace \
--devel
```

Or you can specify a version (remove the `v` letter at the beginning of the release or tag name):

```bash
VERSION="0.2.5-alpha.6"
helm upgrade -i robot-operator robolaunch/robot-operator  \
--namespace robot-system \
--create-namespace \
--version $VERSION
```

### via Manifest

Deploy DevSpace Operator one-file YAML using the command below:

```bash
# select a tag
TAG="v0.2.5-alpha.6"
kubectl apply -f https://raw.githubusercontent.com/robolaunch/robot-operator/$TAG/hack/deploy/manifests/robot_operator.yaml
```

## Uninstalling DevSpace Operator

To uninstall DevSpace Operator installed with Helm, run the following commands:

```bash
helm delete robot-operator -n robot-system
kubectl delete ns robot-system
```

To uninstall DevSpace Operator installed with one-file YAML, run the following commands:
```bash
# find the tag you installed
TAG="v0.2.5-alpha.6"
kubectl delete -f https://raw.githubusercontent.com/robolaunch/robot-operator/$TAG/hack/deploy/manifests/robot_operator.yaml
```