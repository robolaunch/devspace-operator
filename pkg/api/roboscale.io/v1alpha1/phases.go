package v1alpha1

type DevSpacePhase string

const (
	DevSpacePhaseEnvironmentNotFound      DevSpacePhase = "EnvironmentNotFound"
	DevSpacePhaseCreatingEnvironment      DevSpacePhase = "CreatingEnvironment"
	DevSpacePhaseConfiguringEnvironment   DevSpacePhase = "ConfiguringEnvironment"
	DevSpacePhaseCreatingDevelopmentSuite DevSpacePhase = "CreatingDevelopmentSuite"
	DevSpacePhaseConfiguringWorkspaces    DevSpacePhase = "ConfiguringWorkspaces"
	DevSpacePhaseEnvironmentReady         DevSpacePhase = "EnvironmentReady"
	DevSpacePhaseBuilding                 DevSpacePhase = "Building"
	DevSpacePhaseBuilt                    DevSpacePhase = "Built"
	DevSpacePhaseLaunching                DevSpacePhase = "Launching"
	DevSpacePhaseRunning                  DevSpacePhase = "Running"
	DevSpacePhaseDeletingLoaderJob        DevSpacePhase = "DeletingLoaderJob"
	DevSpacePhaseDeletingVolumes          DevSpacePhase = "DeletingVolumes"

	DevSpacePhaseFailed DevSpacePhase = "Failed"
)

type WorkspaceManagerPhase string

const (
	WorkspaceManagerPhaseConfiguringWorkspaces WorkspaceManagerPhase = "ConfiguringWorkspaces"
	WorkspaceManagerPhaseReady                 WorkspaceManagerPhase = "Ready"
	WorkspaceManagerPhaseFailed                WorkspaceManagerPhase = "Failed"
)

type DevSuitePhase string

const (
	DevSuitePhaseDevSpaceNotFound    DevSuitePhase = "DevSpaceNotFound"
	DevSuitePhaseCreatingDevSpaceVDI DevSuitePhase = "CreatingDevSpaceVDI"
	DevSuitePhaseCreatingDevSpaceIDE DevSuitePhase = "CreatingDevSpaceIDE"
	DevSuitePhaseRunning             DevSuitePhase = "Running"
	DevSuitePhaseDeactivating        DevSuitePhase = "Deactivating"
	DevSuitePhaseInactive            DevSuitePhase = "Inactive"
)

type DevSpaceIDEPhase string

const (
	DevSpaceIDEPhaseCreatingService DevSpaceIDEPhase = "CreatingService"
	DevSpaceIDEPhaseCreatingPod     DevSpaceIDEPhase = "CreatingPod"
	DevSpaceIDEPhaseCreatingIngress DevSpaceIDEPhase = "CreatingIngress"
	DevSpaceIDEPhaseRunning         DevSpaceIDEPhase = "Running"
)

type DevSpaceVDIPhase string

const (
	DevSpaceVDIPhaseCreatingPVC        DevSpaceVDIPhase = "CreatingPVC"
	DevSpaceVDIPhaseCreatingTCPService DevSpaceVDIPhase = "CreatingTCPService"
	DevSpaceVDIPhaseCreatingUDPService DevSpaceVDIPhase = "CreatingUDPService"
	DevSpaceVDIPhaseCreatingPod        DevSpaceVDIPhase = "CreatingPod"
	DevSpaceVDIPhaseCreatingIngress    DevSpaceVDIPhase = "CreatingIngress"
	DevSpaceVDIPhaseRunning            DevSpaceVDIPhase = "Running"
)

type DevSpaceJupyterPhase string

const (
	DevSpaceJupyterPhaseCreatingService DevSpaceJupyterPhase = "CreatingService"
	DevSpaceJupyterPhaseCreatingPod     DevSpaceJupyterPhase = "CreatingPod"
	DevSpaceJupyterPhaseCreatingIngress DevSpaceJupyterPhase = "CreatingIngress"
	DevSpaceJupyterPhaseRunning         DevSpaceJupyterPhase = "Running"
)
