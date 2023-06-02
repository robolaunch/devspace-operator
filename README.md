# <img src="https://raw.githubusercontent.com/robolaunch/trademark/main/logos/svg/rocket.svg" width="40" height="40" align="top"> robolaunch Kubernetes DevSpace Operator

<div align="center">
  <p align="center">
    <a href="https://github.com/robolaunch/devspace-operator/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/robolaunch/devspace-operator" alt="license">
    </a>
    <a href="https://github.com/robolaunch/devspace-operator/issues">
      <img src="https://img.shields.io/github/issues/robolaunch/devspace-operator" alt="issues">
    </a>
    <a href="https://github.com/robolaunch/devspace-operator/releases">
      <img src="https://img.shields.io/github/v/release/robolaunch/devspace-operator" alt="release">
    </a>
  </p>
</div>

<div align="center">
  <p align="center">
    <a href="https://github.com/robolaunch/devspace-operator/releases">
      <img src="https://img.shields.io/github/go-mod/go-version/robolaunch/devspace-operator" alt="release">
    </a>
    <a href="https://pkg.go.dev/github.com/robolaunch/devspace-operator">
      <img src="https://pkg.go.dev/badge/github.com/robolaunch/devspace-operator.svg" alt="Go Reference">
    </a>
    <a href="https://goreportcard.com/report/github.com/robolaunch/devspace-operator">
      <img src="https://goreportcard.com/badge/github.com/robolaunch/devspace-operator" alt="Go Reference">
    </a>
  </p>
</div>

<div align="center">
  <p align="center">
    <a href="https://hub.docker.com/u/robolaunchio/devspace-controller-manager">
      <img src="https://img.shields.io/docker/pulls/robolaunchio/devspace-controller-manager" alt="pulls">
    </a>
    <a href="https://github.com/robolaunch/devspace-operator/actions">
      <img src="https://github.com/robolaunch/devspace-operator/actions/workflows/docker-build-for-push.yml/badge.svg" alt="build">
    </a>
  </p>
</div>

robolaunch Kubernetes DevSpace Operator manages lifecycle of Kubernetes native development environments and enables defining, deploying and distributing them declaratively.

<!-- <img src="https://raw.githubusercontent.com/robolaunch/trademark/main/repository-media/devspace-operator/kubectl-get-devspaces.gif" alt="kubectl-get-devspaces" width="100%"/>

<img src="https://raw.githubusercontent.com/robolaunch/trademark/main/repository-media/devspace-operator/kubectl-describe-devspace.gif" alt="kubectl-describe-devspace" width="100%"/> -->

## Table of Contents  
- [Quick Start](#quick-start)
  - [Installation](#installation)
  - [Deploy Your First DevSpace](#deploy-your-first-devspace)
- [Contributing](#contributing)


## Quick Start

### Installation

Label a node in your cluster:

```bash
kubectl label <NODE> robolaunch.io/organization=robolaunch
kubectl label <NODE> robolaunch.io/team=video-processing
kubectl label <NODE> robolaunch.io/region=europe-east
kubectl label <NODE> robolaunch.io/cloud-instance=cluster
kubectl label <NODE> robolaunch.io/cloud-instance-alias=cluster-alias
```

Install DevSpace Operator with Helm:

```bash
# add robolaunch Helm repository and update
helm repo add robolaunch https://robolaunch.github.io/charts/
helm repo update
# install chart
helm upgrade -i devspace-operator robolaunch/devspace-operator  \
--namespace devspace-system \
--create-namespace \
--devel
```

See [installation guide for more](./docs/installation/README.md).

### Deploy Your First DevSpace

DevSpace deployment steps will be instructed here.

## Contributing

Please see [this guide](./CONTRIBUTING) if you want to contribute.
