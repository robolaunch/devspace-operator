<a name="unreleased"></a>
## [Unreleased]


<a name="v0.1.0-alpha.1"></a>
## v0.1.0-alpha.1 - 2023-06-02
### Feat
- **client:** create client for robot
- **discovery-server:** :rocket: manage a discovery server attached to the robot
- **dynamic-ws:** make workspaces changeable dynamically
- **environment:** :rocket: select environment and application through api
- **gpu:** :rocket: watch volatile gpu utilization of node
- **image:** :rocket: select robot image by querying versioning map
- **launch:** :rocket: add custom launch option
- **metrics:** enable multiple ros 2 distribution option
- **metrics:** enable observing robot's metrics
- **metrics-exporter:** follow gpu and network metrics from host
- **network:** :rocket: watch network load of node network interfaces
- **rmw-implementation:** enable selecting rmw implementation
- **robot:** :rocket: manage main robot assets
- **robot-build:** add automation to the building process of robot
- **robot-dev-suite:** manage development suites attached to a robot
- **robot-ide:** provision cloud ide attached to robot
- **robot-launch:** add automation to the launching process of robot
- **robot-vdi:** provision virtual desktop attached to robot
- **ros2-run:** :rocket: support ros2 run command
- **rosbridge:** :rocket: manage ros bridge server for robot instance
- **vdi:** select vdi resolution
- **volume:** :rocket: provision robot's volumes dynamically
- **volumes:** :rocket: configure robot's volumes
- **workspaces:** :rocket: prepare workspaces according to the definitions in robot manifest

### Fix
- **api:** fix rmw implementation types
- **attachments:** add missing case for attachments
- **build-manager:** make phase ready if no step is assigned in instance
- **build-manager:** :bug: fix cluster selection for steps
- **check:** add pvc status check
- **column:** fix column name and key
- **discovery-server:** fix robot with non-attached discovery server
- **fastdds:** :bug: fix fastdds config, use only udpv4
- **field:** change problematic dir names because of fabric8 code generation issues
- **ide:** update ide host and path values in ingress
- **import:** fix broken imports
- **ingress:** :bug: update ingress configurations
- **injections:** fix non-permanent injections to containers
- **ip:** fix broken ip format
- **labels:** :bug: get tenancy labels from robot instead of buildmanager
- **labels:** fix robot image label key
- **launch-manager:** :bug: fix launch pod creation issue
- **launchmanager:** :bug: fix checking irrelevant type equalty in webhooks
- **oauth2:** fix oauth2 url
- **path:** fix api path in project file
- **robot-dev-suite:** :bug: fix syncing component specs
- **robot-dev-suite:** fix checking robot ide
- **robot-ide:** update display connection sources
- **scheduling:** stop assigning status to all steps in deletion attempts
- **status:** clear status from robot
- **status:** fix clearing the workload status if managers are not active
- **steps:** fix instance scheduling of buildmanager steps
- **typo:** fix vdi resource name
- **typo:** fix typo in logs
- **vdi:** update vdi host and path values in ingress
- **vdi:** :bug: fix vdi ingress and connection url

### Pull Requests
- Merge pull request [#24](https://github.com/robolaunch/devspace-operator/issues/24) from robolaunch/23-allow-multiple-launches


[Unreleased]: https://github.com/robolaunch/devspace-operator/compare/v0.1.0-alpha.1...HEAD
