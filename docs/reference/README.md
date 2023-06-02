# API Reference


## dev.roboscale.io/v1alpha1

Package v1alpha1 contains API Schema definitions for the dev v1alpha1 API group

### Resource Types
- [DevSpace](#devspace)
- [WorkspaceManager](#workspacemanager)
- [DevSuite](#devsuite)
- [DevSpaceVDI](#devspacevdi)
- [DevSpaceIDE](#devspaceide)



#### DevSpace



DevSpace is the custom resource that contains ROS 2 components (Workloads, Cloud VDI, Cloud IDE, ROS Bridge, Configurational Resources), robolaunch DevSpace instances can be decomposed and distributed to both cloud instances and physical instances using federation.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `dev.roboscale.io/v1alpha1`
| `kind` _string_ | `DevSpace`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[DevSpaceSpec](#devspacespec)_ | Specification of the desired behavior of the DevSpace. |
| `status` _[DevSpaceStatus](#devspacestatus)_ | Most recently observed status of the DevSpace. |


#### DevSpaceSpec



DevSpaceSpec defines the desired state of DevSpace.

_Appears in:_
- [DevSpace](#devspace)

| Field | Description |
| --- | --- |
| `environment` _[Environment](#environment)_ | Environment properties. Supported options are listed in [robolaunch Platform Versioning Map](https://github.com/robolaunch/robolaunch/blob/main/platform.yaml). |
| `storage` _[Storage](#storage)_ | Total storage amount to persist via DevSpace. Unit of measurement is MB. (eg. `10240` corresponds 10 GB) This amount is being shared between different components. |
| `devSuiteTemplate` _[DevSuiteSpec](#devsuitespec)_ | DevSpace development suite template |
| `workspaceManagerTemplate` _[WorkspaceManagerSpec](#workspacemanagerspec)_ | Workspace manager template to configure ROS 2 workspaces. |
| `development` _boolean_ | [*alpha*] Switch to development mode if `true`. |
| `rootDNSConfig` _[RootDNSConfig](#rootdnsconfig)_ | [*alpha*] Root DNS configuration. |
| `tlsSecretRef` _[TLSSecretReference](#tlssecretreference)_ | [*alpha*] TLS secret reference. |


#### DevSpaceStatus



DevSpaceStatus defines the observed state of DevSpace.

_Appears in:_
- [DevSpace](#devspace)

| Field | Description |
| --- | --- |
| `phase` _[DevSpacePhase](#devspacephase)_ | Phase of DevSpace. It sums the general status of DevSpace. |
| `image` _string_ | Main image of DevSpace. It is derived either from the specifications or determined directly over labels. |
| `nodeName` _string_ | Node that DevSpace uses. It is selected via tenancy labels. |
| `volumeStatuses` _[VolumeStatuses](#volumestatuses)_ | DevSpace persists some of the directories of underlying OS inside persistent volumes. This field exposes persistent volume claims that dynamically provision PVs. |
| `loaderJobStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of loader job that configures environment. |
| `workspaceManagerStatus` _[WorkspaceManagerInstanceStatus](#workspacemanagerinstancestatus)_ | Workspace manager instance status if exists. |
| `devSuiteStatus` _[DevSuiteInstanceStatus](#devsuiteinstancestatus)_ | DevSpace development suite instance status. |
| `attachedDevObjects` _[AttachedDevObject](#attacheddevobject) array_ | [*alpha*] Attached dev object information. |


#### WorkspaceManager



WorkspaceManager configures the ROS 2 workspaces and repositories by executing Kubernetes jobs.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `dev.roboscale.io/v1alpha1`
| `kind` _string_ | `WorkspaceManager`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[WorkspaceManagerSpec](#workspacemanagerspec)_ | Specification of the desired behavior of the WorkspaceManager. |
| `status` _[WorkspaceManagerStatus](#workspacemanagerstatus)_ | Most recently observed status of the WorkspaceManager. |


#### WorkspaceManagerSpec



WorkspaceManagerSpec defines the desired state of WorkspaceManager.

_Appears in:_
- [DevSpaceSpec](#devspacespec)
- [WorkspaceManager](#workspacemanager)

| Field | Description |
| --- | --- |
| `workspacesPath` _string_ | Global path of workspaces. It's fixed to `/root/workspaces` path. |
| `workspaces` _[Workspace](#workspace) array_ | Workspace definitions of devspace. Multiple ROS 2 workspaces can be configured over this field. |
| `updateNeeded` _boolean_ | WorkspaceManager is triggered if this field is set to `true`. Then the workspaces are being configured again while backing up the old configurations. This field is often used by operator. |


#### WorkspaceManagerStatus



WorkspaceManagerStatus defines the observed state of WorkspaceManager.

_Appears in:_
- [WorkspaceManager](#workspacemanager)
- [WorkspaceManagerInstanceStatus](#workspacemanagerinstancestatus)

| Field | Description |
| --- | --- |
| `phase` _[WorkspaceManagerPhase](#workspacemanagerphase)_ | Phase of WorkspaceManager. |
| `clonerJobStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of cloner job. |
| `cleanupJobStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of cleanup jobs that runs while reconfiguring workspaces. |
| `version` _integer_ | Incremental version of workspace configuration map. Used to determine changes in configuration. |


#### DevSuite



DevSuite is a custom resource that creates dynamically configured development environments for devspaces.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `dev.roboscale.io/v1alpha1`
| `kind` _string_ | `DevSuite`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[DevSuiteSpec](#devsuitespec)_ | Specification of the desired behavior of the DevSuite. |
| `status` _[DevSuiteStatus](#devsuitestatus)_ | Most recently observed status of the DevSuite. |


#### DevSuiteSpec



DevSuiteSpec defines the desired state of DevSuite.

_Appears in:_
- [DevSpaceSpec](#devspacespec)
- [DevSuite](#devsuite)

| Field | Description |
| --- | --- |
| `vdiEnabled` _boolean_ | If `true`, a Cloud VDI will be provisioned inside development suite. |
| `devSpaceVDITemplate` _[DevSpaceVDISpec](#devspacevdispec)_ | Configurational parameters of DevSpaceVDI. Applied if `.spec.vdiEnabled` is set to `true`. |
| `ideEnabled` _boolean_ | If `true`, a Cloud IDE will be provisioned inside development suite. |
| `devSpaceIDETemplate` _[DevSpaceIDESpec](#devspaceidespec)_ | Configurational parameters of DevSpaceIDE. Applied if `.spec.ideEnabled` is set to `true`. |


#### DevSuiteStatus



DevSuiteStatus defines the observed state of DevSuite.

_Appears in:_
- [AttachedDevObject](#attacheddevobject)
- [DevSuite](#devsuite)
- [DevSuiteInstanceStatus](#devsuiteinstancestatus)

| Field | Description |
| --- | --- |
| `phase` _[DevSuitePhase](#devsuitephase)_ | Phase of DevSuite. |
| `devSpaceVDIStatus` _[OwnedDevSpaceServiceStatus](#owneddevspaceservicestatus)_ | Status of DevSpaceVDI. |
| `devSpaceIDEStatus` _[OwnedDevSpaceServiceStatus](#owneddevspaceservicestatus)_ | Status of DevSpaceIDE. |
| `active` _boolean_ | [*alpha*] Indicates if DevSuite is attached to a DevSpace and actively provisioned it's resources. |


#### DevSpaceVDI



DevSpaceVDI creates and manages Cloud VDI resources and workloads.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `dev.roboscale.io/v1alpha1`
| `kind` _string_ | `DevSpaceVDI`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[DevSpaceVDISpec](#devspacevdispec)_ | Specification of the desired behavior of the DevSpaceVDI. |
| `status` _[DevSpaceVDIStatus](#devspacevdistatus)_ | Most recently observed status of the DevSpaceVDI. |


#### DevSpaceVDISpec



DevSpaceVDISpec defines the desired state of DevSpaceVDI.

_Appears in:_
- [DevSpaceVDI](#devspacevdi)
- [DevSuiteSpec](#devsuitespec)

| Field | Description |
| --- | --- |
| `resources` _[Resources](#resources)_ | Resource limitations of Cloud IDE. |
| `serviceType` _[ServiceType](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#servicetype-v1-core)_ | Service type of Cloud IDE. `ClusterIP` and `NodePort` is supported. |
| `privileged` _boolean_ | If `true`, containers of DevSpaceIDE will be privileged containers. It can be used in physical instances where it's necessary to access I/O devices on the host machine. Not recommended to activate this field on cloud instances. |
| `nat1to1` _string_ | NAT1TO1 option for Cloud VDI. |
| `webrtcPortRange` _string_ | UDP port range to used in WebRTC connections. |
| `resolution` _string_ | VDI screen resolution options. Default is `2048x1152`. |
| `ingress` _boolean_ | [*alpha*] DevSpaceIDE will create an Ingress resource if `true`. |


#### DevSpaceVDIStatus



DevSpaceVDIStatus defines the observed state of DevSpaceVDI.

_Appears in:_
- [DevSpaceVDI](#devspacevdi)

| Field | Description |
| --- | --- |
| `phase` _DevSpaceVDIPhase_ | Phase of DevSpaceVDI. |
| `podStatus` _[OwnedPodStatus](#ownedpodstatus)_ | Status of Cloud VDI pod. |
| `serviceTCPStatus` _[OwnedServiceStatus](#ownedservicestatus)_ | Status of Cloud VDI TCP service. |
| `serviceUDPStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud VDI UDP service. |
| `ingressStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud VDI Ingress. |
| `pvcStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud VDI persistent volume claim. This PVC dynamically provisions a volume that is a shared between DevSpaceVDI workloads and other workloads that requests display. |


#### DevSpaceIDE



DevSpaceIDE creates and manages Cloud IDE resources and workloads.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `dev.roboscale.io/v1alpha1`
| `kind` _string_ | `DevSpaceIDE`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[DevSpaceIDESpec](#devspaceidespec)_ | Specification of the desired behavior of the DevSpaceIDE. |
| `status` _[DevSpaceIDEStatus](#devspaceidestatus)_ | Most recently observed status of the DevSpaceIDE. |


#### DevSpaceIDESpec



DevSpaceIDESpec defines the desired state of DevSpaceIDE.

_Appears in:_
- [DevSpaceIDE](#devspaceide)
- [DevSuiteSpec](#devsuitespec)

| Field | Description |
| --- | --- |
| `resources` _[Resources](#resources)_ | Resource limitations of Cloud IDE. |
| `serviceType` _[ServiceType](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#servicetype-v1-core)_ | Service type of Cloud IDE. `ClusterIP` and `NodePort` is supported. |
| `privileged` _boolean_ | If `true`, containers of DevSpaceIDE will be privileged containers. It can be used in physical instances where it's necessary to access I/O devices on the host machine. Not recommended to activate this field on cloud instances. |
| `display` _boolean_ | Cloud IDE connects an X11 socket if it's set to `true` and a target DevSpaceVDI resource is set in labels with key `robolaunch.io/target-vdi`. Applications that requires GUI can be executed such as rViz. |
| `ingress` _boolean_ | [*alpha*] DevSpaceIDE will create an Ingress resource if `true`. |


#### DevSpaceIDEStatus



DevSpaceIDEStatus defines the observed state of DevSpaceIDE.

_Appears in:_
- [DevSpaceIDE](#devspaceide)

| Field | Description |
| --- | --- |
| `phase` _DevSpaceIDEPhase_ | Phase of DevSpaceIDE. |
| `podStatus` _[OwnedPodStatus](#ownedpodstatus)_ | Status of Cloud IDE pod. |
| `serviceStatus` _[OwnedServiceStatus](#ownedservicestatus)_ | Status of Cloud IDE service. |
| `ingressStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud IDE Ingress. |


#### Application





_Appears in:_
- [Environment](#environment)

| Field | Description |
| --- | --- |
| `name` _string_ | Application name. |
| `version` _string_ | Version of the application. |


#### AttachedDevObject





_Appears in:_
- [DevSpaceStatus](#devspacestatus)

| Field | Description |
| --- | --- |
| `reference` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectreference-v1-core)_ | Reference to the DevSuite instance. |
| `status` _[DevSuiteStatus](#devsuitestatus)_ | Status of attached DevSuite. |


#### DevSpaceImage





_Appears in:_
- [Environment](#environment)

| Field | Description |
| --- | --- |
| `ubuntuDistro` _string_ | Ubuntu distribution of the environment. |
| `desktop` _string_ | Ubuntu desktop. |
| `version` _string_ | DevSpace image version. |


#### DevSpacePhase

_Underlying type:_ `string`



_Appears in:_
- [DevSpaceStatus](#devspacestatus)



#### DevSuiteInstanceStatus





_Appears in:_
- [DevSpaceStatus](#devspacestatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `status` _[DevSuiteStatus](#devsuitestatus)_ | Status of the DevSuite instance. |


#### DevSuitePhase

_Underlying type:_ `string`



_Appears in:_
- [DevSuiteStatus](#devsuitestatus)



#### Environment





_Appears in:_
- [DevSpaceSpec](#devspacespec)

| Field | Description |
| --- | --- |
| `domain` _string_ | Domain of the environment. |
| `application` _[Application](#application)_ | Application properties. |
| `devspace` _[DevSpaceImage](#devspaceimage)_ | DevSpace image properties. |


#### OwnedDevSpaceServiceStatus





_Appears in:_
- [DevSuiteStatus](#devsuitestatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `connection` _string_ | Address of the devspace service that can be reached from outside. |


#### OwnedPodStatus





_Appears in:_
- [DevSpaceIDEStatus](#devspaceidestatus)
- [DevSpaceVDIStatus](#devspacevdistatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `ip` _string_ | IP of the pod. |


#### OwnedResourceStatus



Generic status for any owned resource.

_Appears in:_
- [DevSpaceIDEStatus](#devspaceidestatus)
- [DevSpaceStatus](#devspacestatus)
- [DevSpaceVDIStatus](#devspacevdistatus)
- [DevSuiteInstanceStatus](#devsuiteinstancestatus)
- [OwnedDevSpaceServiceStatus](#owneddevspaceservicestatus)
- [OwnedPodStatus](#ownedpodstatus)
- [OwnedServiceStatus](#ownedservicestatus)
- [VolumeStatuses](#volumestatuses)
- [WorkspaceManagerInstanceStatus](#workspacemanagerinstancestatus)
- [WorkspaceManagerStatus](#workspacemanagerstatus)

| Field | Description |
| --- | --- |
| `created` _boolean_ | Shows if the owned resource is created. |
| `reference` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectreference-v1-core)_ | Reference to the owned resource. |
| `phase` _string_ | Phase of the owned resource. |


#### OwnedServiceStatus





_Appears in:_
- [DevSpaceIDEStatus](#devspaceidestatus)
- [DevSpaceVDIStatus](#devspacevdistatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `url` _string_ | Connection URL. |


#### Repository



Repository description.

_Appears in:_
- [Workspace](#workspace)

| Field | Description |
| --- | --- |
| `url` _string_ | Base URL of the repository. |
| `branch` _string_ | Branch of the repository to clone. |
| `path` _string_ | [*Autofilled*] Absolute path of repository |
| `owner` _string_ | [*Autofilled*] User or organization, maintainer of repository |
| `repo` _string_ | [*Autofilled*] Repository name |
| `hash` _string_ | [*Autofilled*] Hash of last commit |


#### Resources



VDI resource limits.

_Appears in:_
- [DevSpaceIDESpec](#devspaceidespec)
- [DevSpaceVDISpec](#devspacevdispec)

| Field | Description |
| --- | --- |
| `gpuCore` _integer_ | GPU core number that will be allocated. |
| `cpu` _string_ | CPU resource limit. |
| `memory` _string_ | Memory resource limit. |


#### RootDNSConfig





_Appears in:_
- [DevSpaceSpec](#devspacespec)

| Field | Description |
| --- | --- |
| `host` _string_ | [*alpha*] Root DNS name.. |


#### Storage



DevSpace's resource limitations.

_Appears in:_
- [DevSpaceSpec](#devspacespec)

| Field | Description |
| --- | --- |
| `amount` _integer_ | Specifies how much storage will be allocated in total. Use MB as a unit of measurement. (eg. `10240` is equal to 10 GB) |
| `storageClassConfig` _[StorageClassConfig](#storageclassconfig)_ | Storage class selection for devspace's volumes. |


#### StorageClassConfig



Storage class configuration for a volume type.

_Appears in:_
- [Storage](#storage)

| Field | Description |
| --- | --- |
| `name` _string_ | Storage class name. |
| `accessMode` _[PersistentVolumeAccessMode](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#persistentvolumeaccessmode-v1-core)_ | PVC access modes. Currently, only `ReadWriteOnce` is supported. |


#### TLSSecretReference





_Appears in:_
- [DevSpaceSpec](#devspacespec)

| Field | Description |
| --- | --- |
| `name` _string_ | [*alpha*] TLS secret object name. |
| `namespace` _string_ | [*alpha*] TLS secret object namespace. |


#### VolumeStatuses





_Appears in:_
- [DevSpaceStatus](#devspacestatus)

| Field | Description |
| --- | --- |
| `varDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/var` directory of underlying OS. |
| `etcDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/etc` directory of underlying OS. |
| `usrDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/usr` directory of underlying OS. |
| `optDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/opt` directory of underlying OS. |
| `workspaceDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the workspaces directory of underlying OS. |


#### Workspace



Workspace description. Each devspace should contain at least one workspace. A workspace should contain at least one repository in it.

_Appears in:_
- [WorkspaceManagerSpec](#workspacemanagerspec)

| Field | Description |
| --- | --- |
| `name` _string_ | Name of workspace. If a workspace's name is `my_ws`, it's absolute path is `/home/workspaces/my_ws`. |
| `repositories` _object (keys:string, values:[Repository](#repository))_ | Repositories to clone inside workspace's `src` directory. |


#### WorkspaceManagerInstanceStatus





_Appears in:_
- [DevSpaceStatus](#devspacestatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `status` _[WorkspaceManagerStatus](#workspacemanagerstatus)_ | Status of the WorkspaceManager instance. |


#### WorkspaceManagerPhase

_Underlying type:_ `string`



_Appears in:_
- [WorkspaceManagerStatus](#workspacemanagerstatus)



