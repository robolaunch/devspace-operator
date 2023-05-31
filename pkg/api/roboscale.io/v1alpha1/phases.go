package v1alpha1

type DevspacePhase string

const (
	DevspacePhaseCreatingEnvironment      DevspacePhase = "CreatingEnvironment"
	DevspacePhaseConfiguringEnvironment   DevspacePhase = "ConfiguringEnvironment"
	DevspacePhaseCreatingDevelopmentSuite DevspacePhase = "CreatingDevelopmentSuite"
	DevspacePhaseConfiguringWorkspaces    DevspacePhase = "ConfiguringWorkspaces"
	DevspacePhaseEnvironmentReady         DevspacePhase = "EnvironmentReady"
	DevspacePhaseBuilding                 DevspacePhase = "Building"
	DevspacePhaseBuilt                    DevspacePhase = "Built"
	DevspacePhaseLaunching                DevspacePhase = "Launching"
	DevspacePhaseRunning                  DevspacePhase = "Running"
	DevspacePhaseDeletingLoaderJob        DevspacePhase = "DeletingLoaderJob"
	DevspacePhaseDeletingVolumes          DevspacePhase = "DeletingVolumes"

	DevspacePhaseFailed DevspacePhase = "Failed"
)

type WorkspaceManagerPhase string

const (
	WorkspaceManagerPhaseConfiguringWorkspaces WorkspaceManagerPhase = "ConfiguringWorkspaces"
	WorkspaceManagerPhaseReady                 WorkspaceManagerPhase = "Ready"
	WorkspaceManagerPhaseFailed                WorkspaceManagerPhase = "Failed"
)

type DevSuitePhase string

const (
	DevSuitePhaseDevspaceNotFound    DevSuitePhase = "DevspaceNotFound"
	DevSuitePhaseCreatingDevSpaceVDI DevSuitePhase = "CreatingDevSpaceVDI"
	DevSuitePhaseCreatingDevSpaceIDE DevSuitePhase = "CreatingDevspaceIDE"
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
