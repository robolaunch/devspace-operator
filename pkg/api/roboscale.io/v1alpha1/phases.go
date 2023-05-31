package v1alpha1

type RobotPhase string

const (
	RobotPhaseCreatingEnvironment      RobotPhase = "CreatingEnvironment"
	RobotPhaseConfiguringEnvironment   RobotPhase = "ConfiguringEnvironment"
	RobotPhaseCreatingDevelopmentSuite RobotPhase = "CreatingDevelopmentSuite"
	RobotPhaseConfiguringWorkspaces    RobotPhase = "ConfiguringWorkspaces"
	RobotPhaseEnvironmentReady         RobotPhase = "EnvironmentReady"
	RobotPhaseBuilding                 RobotPhase = "Building"
	RobotPhaseBuilt                    RobotPhase = "Built"
	RobotPhaseLaunching                RobotPhase = "Launching"
	RobotPhaseRunning                  RobotPhase = "Running"
	RobotPhaseDeletingLoaderJob        RobotPhase = "DeletingLoaderJob"
	RobotPhaseDeletingVolumes          RobotPhase = "DeletingVolumes"

	RobotPhaseFailed RobotPhase = "Failed"
)

type WorkspaceManagerPhase string

const (
	WorkspaceManagerPhaseConfiguringWorkspaces WorkspaceManagerPhase = "ConfiguringWorkspaces"
	WorkspaceManagerPhaseReady                 WorkspaceManagerPhase = "Ready"
	WorkspaceManagerPhaseFailed                WorkspaceManagerPhase = "Failed"
)

type DevSuitePhase string

const (
	DevSuitePhaseRobotNotFound       DevSuitePhase = "RobotNotFound"
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
